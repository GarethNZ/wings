# steamcmd Base Installation Script
#
# Server Files: /mnt/server
# Image to install with is 'windows nano server'

Set-PSDebug -Trace 1

##
#
# Variables
# STEAM_USER, STEAM_PASS, STEAM_AUTH - Steam user setup. If a user has 2fa enabled it will most likely fail due to timeout. Leave blank for anon install.
# SRCDS_APPID - steam app id found here - https://developer.valvesoftware.com/wiki/Dedicated_Servers_List
# EXTRA_FLAGS - when a server has extra flags for things like beta installs or updates.
#
##

## just in case someone removed the defaults.
If ( [string]::IsNullOrWhiteSpace($ENV:STEAM_USER) ) {
    Write-Output "steam user is not set.\n"
    Write-Output "Using anonymous user.\n"
    $ENV:STEAM_USER="anonymous"
    $ENV:STEAM_PASS=""
    $ENV:STEAM_AUTH=""
}
else {
    Write-Output "user set to $ENV:STEAM_USER"
}

## download and install steamcmd
$ENV:HOME="C:/mnt/server"

mkdir -p $ENV:HOME/steamcmd
Invoke-WebRequest -Uri https://steamcdn-a.akamaihd.net/client/installer/steamcmd.zip -OutFile $ENV:HOME/steamcmd.zip
Expand-Archive -Path $ENV:HOME/steamcmd.zip -DestinationPath $ENV:HOME/steamcmd -Force
Set-Location $ENV:HOME/steamcmd

## install game using steamcmd
./steamcmd.exe +force_install_dir $ENV:HOME +login $ENV:STEAM_USER $ENV:STEAM_PASS $ENV:STEAM_AUTH +app_update $ENV:SRCDS_APPID $ENV:EXTRA_FLAGS +quit ## other flags may be needed depending on install. looking at you cs 1.6
