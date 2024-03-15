 #!bin/bash

 ipServer=$2
 nameUser=$1
 patch=$3

 echo "Start script $0"


 antifraud=$(ssh "$nameUser"@"$ipServer" ps aux 2>/dev/null | grep -wo [a]ntifraud);

 if [ "$antifraud" == "antifraud" ]; then
 	echo "Killing antifraud process";
 	ssh "$nameUser"@"$ipServer" "kill $(ssh "$nameUser"@"$ipServer" ps aux 2>/dev/null | grep '[a]ntifraud' | awk '{print $2}')";
 	echo "Removed old antifraud "
 	ssh "$nameUser"@"$ipServer" "rm antifraud";
 fi

 scp  "$patch" "$nameUser"@"$ipServer":.;

 ssh "$nameUser"@"$ipServer" "./antifraud";