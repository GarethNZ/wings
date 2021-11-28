Set-PSDebug -Trace 1

Set-Location /home/container
Start-Sleep 1

# Make internal Docker IP address available to processes.
# Not Available to windows containers as it uses NAT
# export INTERNAL_IP=`ip route get 1 | awk '{print $NF;exit}'`

# Update Source Server
If ( ![string]::IsNullOrWhiteSpace($ENV:SRCDS_APPID) ) {
    If ( ![string]::IsNullOrWhiteSpace($ENV:SRCDS_BETAID) ) {
        If ( ![string]::IsNullOrWhiteSpace($ENV:SRCDS_BETAPASS) ) {
            ./steamcmd/steamcmd.exe +force_install_dir /home/container +login anonymous +app_update $ENV:SRCDS_APPID -beta $ENV:SRCDS_BETAID -betapassword $ENV:SRCDS_BETAPASS +quit
        }
        else {
            ./steamcmd/steamcmd.exe +force_install_dir /home/container +login anonymous +app_update $ENV:SRCDS_APPID -beta $ENV:SRCDS_BETAID +quit
        }
    }
    else {
        ./steamcmd/steamcmd.exe +force_install_dir /home/container +login anonymous +app_update $ENV:SRCDS_APPID +quit
    }
}

# Replace Startup Variables
$MODIFIED_STARTUP = $ENV:STARTUP -replace "{{","%" -replace "}}", "%" | ForEach-Object { [Environment]::ExpandEnvironmentVariables($_) }

# Win Executables spawn their own process and do not block this process
# Assuming they all need this for now
$MODIFIED_STARTUP += " | Wait-Process"

Write-Output ":/home/container$ $MODIFIED_STARTUP"

# Run the Server
Invoke-Expression $MODIFIED_STARTUP

