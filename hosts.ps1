# Get script location for explicit reference

function GetHostFile($HostFileLocation){
  $content = Get-Content $HostFileLocation
  $currentJson = "{$content}"

  return ConvertFrom-Json $currentJson
}

function GetAllHostFiles($Location){
  $json = @{}
  Get-ChildItem -Path $Location | ForEach-Object {
    $jsonHostfile = GetHostFile($_)
    $property = $jsonHostfile.PsObject.Properties
    $json[$property.Name] = $jsonHostfile[$property.Name]
  }
  return $json
}

if($args[1] == '--host'){
  Write-Output '{}'
  exit 0
}

$addresses = Invoke-Expression 'nmap --min-parallelism 50 --min-rate 50 -p 22 10.100.101-103.50-250 --open| grep -v "10.100.101.223" | grep -Eo "[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}"'

$startOut = GetAllHostFiles(Join-Path -Path $PSScriptRoot -ChildPath "hostfiles")

$rack = Invoke-Expression "ip -f inet address | grep 10.100 | cut -d. -f 3"

[System.Collections.ArrayList]$npfGameServers
[System.Collections.ArrayList]$npfGameServersCurrentRack
if([string]::IsNullOrEmpty($addresses)){
  foreach ($ip in $result.Split("`n")) {
    $npfGameServers.Add($ip)

    $ipSlices = [string[]] $ip.split('.')

    if($ipSlices[3] == $rack){
      $npfGameServersCurrentRack.Add($ip)
    }
  }
}
