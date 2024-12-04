# Save this as MeasureCommand.ps1

param (
    [Parameter(Mandatory = $true)]
    [int]$n,

    [Parameter(Mandatory = $true)]
    [string]$cmd
)

# Array to store durations
$durations = @()

# Check for Python or Node.js availability
function Get-Timer {
    if (Get-Command python3 -ErrorAction SilentlyContinue) {
        return "python3"
    } elseif (Get-Command node -ErrorAction SilentlyContinue) {
        return "node"
    } else {
        Write-Error "Python3 or Node.js is required but not found."
        exit 1
    }
}

# Get current time in nanoseconds using available language
function Get-TimeNs {
    param (
        [string]$cmd
    )

    switch ($timer) {
        "python3" {
            $script = @"
import time
import subprocess
import sys
cmd = sys.argv[1]
start = time.perf_counter_ns()
subprocess.run(cmd, shell=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
end = time.perf_counter_ns()
print(end - start)
"@
            python3 -c $script $cmd
        }
        "node" {
            $script = @"
const { execSync } = require('child_process');
const cmd = process.argv[1];
const start = process.hrtime.bigint();
execSync(cmd, {stdio: 'ignore'});
const end = process.hrtime.bigint();
console.log(Number(end - start));
"@
            node -e $script $cmd
        }
        default {
            Write-Error "No suitable language found for nanosecond timing!"
            exit 1
        }
    }
}

$timer = Get-Timer
Write-Host "Using $timer for nanosecond precision timing."
Write-Host "Running command '$cmd' $n times..."

# Run the command n times and measure time
for ($i = 1; $i -le $n; $i++) {
    $duration = Get-TimeNs -cmd $cmd
    if (-not $duration) {
        Write-Error "Error running command. Exiting."
        exit 1
    }
    $durations += [int64]$duration
}

# Calculate statistics
$total = ($durations | Measure-Object -Sum).Sum
$min = ($durations | Measure-Object -Minimum).Minimum
$max = ($durations | Measure-Object -Maximum).Maximum
$avg = [int64]($total / $n)

# Sort durations
$sorted = $durations | Sort-Object

$p90Index = [int]($n * 0.9) - 1
$p95Index = [int]($n * 0.95) - 1
$p99Index = [int]($n * 0.99) - 1

$p90 = $sorted[$p90Index]
$p95 = $sorted[$p95Index]
$p99 = $sorted[$p99Index]

# Convert nanoseconds to milliseconds for display
function Convert-NsToMs($ns) {
    return [Math]::Round($ns / 1e6, 3)
}

Write-Host "Statistics:"
Write-Host "  Total time: $(Convert-NsToMs $total) ms"
Write-Host "  Min time: $(Convert-NsToMs $min) ms"
Write-Host "  Max time: $(Convert-NsToMs $max) ms"
Write-Host "  Avg time: $(Convert-NsToMs $avg) ms"
Write-Host "  P90: $(Convert-NsToMs $p90) ms"
Write-Host "  P95: $(Convert-NsToMs $p95) ms"
Write-Host "  P99: $(Convert-NsToMs $p99) ms"
