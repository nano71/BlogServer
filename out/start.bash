killall blogServer_linux
rm -rf ./nohup.out
nohup ./blogServer_linux -releaseMode=false -protocol=https &
