# ----------------------------------
# Pterodactyl Core Dockerfile
# Environment: Source Engine
# Minimum Panel Version: 0.6.0
# ----------------------------------
#FROM        mcr.microsoft.com/windows/nanoserver:10.0.19042.1348
#FROM        mcr.microsoft.com/powershell:lts-nanoserver-1909 # NOTE: steamcmd does nothing on this? no output at all??? (it is MUCH smaller than server core tho)
FROM        mcr.microsoft.com/powershell:lts-windowsservercore-1909

LABEL       author="Pterodactyl Software" maintainer="gareth@garethnz.com"

ENV         HOME C:/home/container
WORKDIR     C:/home/container

COPY        ./entrypoint.ps1 /entrypoint.ps1
COPY        ./installscript.ps1 /installscript.ps1
CMD         ["powershell", "C:/entrypoint.ps1"]