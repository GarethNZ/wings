$Env:TZ="Pacific/Auckland"
mkdir -p "/usr/lib"
"ID=""distroless""" > /usr/lib/os-release
$ENV:WINGS_USERNAME="user"
# UID and GID irrelevant for windows
$ENV:WINGS_UID="1000"
$ENV:WINGS_GID="1000"

./wings.exe --debug --config=config.yml


# config.yml notes:
# Important parts:
# docker:
#   network:
#     name: nat
#     driver: nat
#     network_mode: nat