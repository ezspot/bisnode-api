# Create all necessary directories
$directories = @(
    "cmd/api",
    "internal/handlers",
    "internal/services/bisnode",
    "internal/models",
    "internal/middleware",
    "internal/response",
    "internal/routes",
    "internal/config"
)

foreach ($dir in $directories) {
    $fullPath = Join-Path -Path $PWD.Path -ChildPath $dir
    if (-not (Test-Path -Path $fullPath)) {
        New-Item -ItemType Directory -Path $fullPath -Force | Out-Null
        Write-Host "Created directory: $fullPath"
    }
}