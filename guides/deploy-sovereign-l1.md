# How to Deploy a Sovereign L1 on the Etna Devnet

Use the CLI to create, deploy, and convert your L1 tracked by a locally run Node.

Warning: this flow is in active development. None of the following should be used in or with production-related infrastructure.

## Build the POC tag of AvalancheGo

```zsh
mkdir -p $GOPATH/src/github.com/ava-labs
cd $GOPATH/src/github.com/ava-labs
git clone https://github.com/ava-labs/avalanchego.git
cd $GOPATH/src/github.com/ava-labs/avalanchego
git checkout v1.12.0-initial-poc.5
./scripts/build.sh
```

## Start your Node and Get Your Node's Credentials

```zsh
➜ avalanchego git:(v1.12.0-initial-poc.5) ./build/avalanchego \
    --network-id="network-76" \
    --bootstrap-ids="NodeID-8LbTmmGsDC991SbD8Nkx88VULT3XYzYXC,NodeID-bojBKDrpt81bYhxYKQfLw89V7CpoH2m7,NodeID-WrLWMK5sJ4dBUAsx1dP2FUyTqrYwbFA1,NodeID-DDhXtFm6Q9tCq2yiFRmcSMKvHgUgh8yQC,NodeID-QDYnWDQd6g4cQ5H6yiWNqSmfRMBqEH9AG" \
    --bootstrap-ips="52.201.126.172:9651,34.233.248.130:9651,107.21.11.213:9651,35.170.144.5:9651,98.82.41.186:9651" \
    --upgrade-file-content="ewogICAgImFwcmljb3RQaGFzZTFUaW1lIjogIjIwMjAtMTItMDVUMDU6MDA6MDBaIiwKICAgICJhcHJpY290UGhhc2UyVGltZSI6ICIyMDIwLTEyLTA1VDA1OjAwOjAwWiIsCiAgICAiYXByaWNvdFBoYXNlM1RpbWUiOiAiMjAyMC0xMi0wNVQwNTowMDowMFoiLAogICAgImFwcmljb3RQaGFzZTRUaW1lIjogIjIwMjAtMTItMDVUMDU6MDA6MDBaIiwKICAgICJhcHJpY290UGhhc2U0TWluUENoYWluSGVpZ2h0IjogMCwKICAgICJhcHJpY290UGhhc2U1VGltZSI6ICIyMDIwLTEyLTA1VDA1OjAwOjAwWiIsCiAgICAiYXByaWNvdFBoYXNlUHJlNlRpbWUiOiAiMjAyMC0xMi0wNVQwNTowMDowMFoiLAogICAgImFwcmljb3RQaGFzZTZUaW1lIjogIjIwMjAtMTItMDVUMDU6MDA6MDBaIiwKICAgICJhcHJpY290UGhhc2VQb3N0NlRpbWUiOiAiMjAyMC0xMi0wNVQwNTowMDowMFoiLAogICAgImJhbmZmVGltZSI6ICIyMDIwLTEyLTA1VDA1OjAwOjAwWiIsCiAgICAiY29ydGluYVRpbWUiOiAiMjAyMC0xMi0wNVQwNTowMDowMFoiLAogICAgImNvcnRpbmFYQ2hhaW5TdG9wVmVydGV4SUQiOiAiMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTFMcG9ZWSIsCiAgICAiZHVyYW5nb1RpbWUiOiAiMjAyMC0xMi0wNVQwNTowMDowMFoiLAogICAgImV0bmFUaW1lIjogIjIwMjQtMTAtMDlUMjA6MDA6MDBaIgp9Cg==" \
    --genesis-file-content="ewogICJuZXR3b3JrSUQiOiA3NiwKICAiYWxsb2NhdGlvbnMiOiBbCiAgICB7CiAgICAgICJldGhBZGRyIjogIjB4QzcxQTYxYTgxNWU0OWQxNkM0MjU0ODJBMzQyYTM2N0NENDJFMzhhNiIsCiAgICAgICJhdmF4QWRkciI6ICJYLWN1c3RvbTF2NnZ1d3hqZ3IwNDNzZzBudXVocTcwazZ2Z251bGU2OTJmdm5yOSIsCiAgICAgICJpbml0aWFsQW1vdW50IjogNTAwMDAwMDAwMDAwMDAwMDAwLAogICAgICAidW5sb2NrU2NoZWR1bGUiOiBbCiAgICAgICAgewogICAgICAgICAgImFtb3VudCI6IDEwMDAwMDAwMDAwMDAwMDAwMCwKICAgICAgICAgICJsb2NrdGltZSI6IDE2MzM4MjQwMDAKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJhbW91bnQiOiAxMDAwMDAwMDAwMDAwMDAwMDAsCiAgICAgICAgICAibG9ja3RpbWUiOiAxNjMzODI1MDAwCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAiYW1vdW50IjogMTAwMDAwMDAwMDAwMDAwMDAwLAogICAgICAgICAgImxvY2t0aW1lIjogMTYzMzgyNjAwMAogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgImFtb3VudCI6IDEwMDAwMDAwMDAwMDAwMDAwMCwKICAgICAgICAgICJsb2NrdGltZSI6IDE2MzM4MjcwMDAKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJhbW91bnQiOiAxMDAwMDAwMDAwMDAwMDAwMDAsCiAgICAgICAgICAibG9ja3RpbWUiOiAxNjMzODI4MDAwCiAgICAgICAgfQogICAgICBdCiAgICB9CiAgXSwKICAic3RhcnRUaW1lIjogMTcyNTMwMDAwMCwKICAiaW5pdGlhbFN0YWtlRHVyYXRpb24iOiAzMTUzMDAwMCwKICAiaW5pdGlhbFN0YWtlRHVyYXRpb25PZmZzZXQiOiA1NDAwLAogICJpbml0aWFsU3Rha2VkRnVuZHMiOiBbCiAgICAiWC1jdXN0b20xdjZ2dXd4amdyMDQzc2cwbnV1aHE3MGs2dmdudWxlNjkyZnZucjkiCiAgXSwKICAiaW5pdGlhbFN0YWtlcnMiOiBbCiAgICB7CiAgICAgICJub2RlSUQiOiAiTm9kZUlELWdwWFdCRXhRU1pYcUpQUXQ2TDZNbnZlVWZncjdISjRxIiwKICAgICAgInJld2FyZEFkZHJlc3MiOiAiWC1jdXN0b20xdjZ2dXd4amdyMDQzc2cwbnV1aHE3MGs2dmdudWxlNjkyZnZucjkiLAogICAgICAiZGVsZWdhdGlvbkZlZSI6IDYyNTAwLAogICAgICAic2lnbmVyIjogewogICAgICAgICJwdWJsaWNLZXkiOiAiMHhhMTRkNjdmMDk3ZDdlNjUxNDY5NmZkODMwODA3OTRiNmI1Y2E2NjQwMDFmMmVkZTRmZDZmMDFkYTQ5MzNkYjg3NWZmMDI4ZmVjNDJiMjlmYzU1MjQ5NDFlMGYyMDgzMGYiLAogICAgICAgICJwcm9vZk9mUG9zc2Vzc2lvbiI6ICIweDgyMzUyYWUxZTAxMDM4MTczZTkyZTA4OGJkMzRjMmJlZTljYzRiMzRkZjVjNWU4YmQyNzczY2VmOTIzOGVlZjg3MGMyZjkzZmE4OTYwNzMzMmNjYmI4NGFhNjY2MDhjNzA2YjdjMmYxMjdiOGI4MGM0NjFjMDRiYmM2MDgyYWZiZmZlMjIwYWFjNzlmNjY1MzNlYTdjNjNmMDQ1MWQ3ZDMyNDU2MzY5ZGQzMzVjOTcxMDkzOGVlNDExMWQwOGQ3OSIKICAgICAgfQogICAgfSwKICAgIHsKICAgICAgIm5vZGVJRCI6ICJOb2RlSUQtNzhpYldwanRaejVaR1Q2RXlURWR1OFZLbWJvVUhUdUdUIiwKICAgICAgInJld2FyZEFkZHJlc3MiOiAiWC1jdXN0b20xdjZ2dXd4amdyMDQzc2cwbnV1aHE3MGs2dmdudWxlNjkyZnZucjkiLAogICAgICAiZGVsZWdhdGlvbkZlZSI6IDYyNTAwLAogICAgICAic2lnbmVyIjogewogICAgICAgICJwdWJsaWNLZXkiOiAiMHg4MzI3ZGJlMWJhNDExYzI3MDYzN2IwODBhODQ3MWZiNDFlZWI4YTliMzkxN2FmMDcyNzUwMWVmOGJkYWE5MDFkMDYzNzgwYmQ3MDJmMzBmNDU4YTYxZjNkNDI5N2RjOTgiLAogICAgICAgICJwcm9vZk9mUG9zc2Vzc2lvbiI6ICIweGE5YzAzOWI1NzY1YWIwNjhiZDYzMmJiY2RjOWJjMmE1M2YyOWUyYzU2YjMzZTMwZDczMmEyM2Q4YzQzMGQ1M2VmNDdlYmNjZmFhNWNmY2VkZDhmMDQxYzJjMTM0OGYwYjBlYWM0MTMxOTJiNzU0NGQyODRmODJkMWZhMGY3NGY5OGQ1ODA1OTA1MzYzYjgxODZlZmRlZjZlNzcxODJmYjFlNzE0N2Y4NTExZTkwMGQxOTVkYjA2ZGE2YTIyZjBhMCIKICAgICAgfQogICAgfSwKICAgIHsKICAgICAgIm5vZGVJRCI6ICJOb2RlSUQtTDRDWThCNXVWU0RlNGNuTjFCcGVEc0hhY01wNHE0cThxIiwKICAgICAgInJld2FyZEFkZHJlc3MiOiAiWC1jdXN0b20xdjZ2dXd4amdyMDQzc2cwbnV1aHE3MGs2dmdudWxlNjkyZnZucjkiLAogICAgICAiZGVsZWdhdGlvbkZlZSI6IDYyNTAwLAogICAgICAic2lnbmVyIjogewogICAgICAgICJwdWJsaWNLZXkiOiAiMHhhOThjNjQ2YThjODYyZWMxNTMyNmU0Y2ZlMmEwZjY2YThmYjdjZjU1NTc2NWY4M2ZmMzIwYTFhNzYyNjgyMjhmM2M4YjI2MmQxZGU0MDA4ZTBiYTQ5YTg5Y2ZhYmZiOTUiLAogICAgICAgICJwcm9vZk9mUG9zc2Vzc2lvbiI6ICIweDk1YzgxNmE0ZDI5MmE0N2M0ZDk5MzRlMzU4NjIxZDA3ODVmZDI5MjBhMWEzMDRjZWJjOWI3ZTQ3NDE3ZTNmZmY3ODBjZmNkZGY3Y2ExMTc3YjQ1YmJmYWMyZjk5Nzg1ODE3NjFkOWRkZDU1ZWM2MTQyZDkyOTk4ZWVhZGJhZmU4Y2Q3NjUxMDU2ZmJiNzlhZmVhNjQzZjBjZDIwZmY0ZjYzODlkZGQ5MWVlMmRiNDU3OTQzOGE2OTA4NjA5YjRjMSIKICAgICAgfQogICAgfSwKICAgIHsKICAgICAgIm5vZGVJRCI6ICJOb2RlSUQtUDVRR0g0RVhkZHJjeU5BemtxeVpLSFhnRXBWWDZIRXhMIiwKICAgICAgInJld2FyZEFkZHJlc3MiOiAiWC1jdXN0b20xdjZ2dXd4amdyMDQzc2cwbnV1aHE3MGs2dmdudWxlNjkyZnZucjkiLAogICAgICAiZGVsZWdhdGlvbkZlZSI6IDYyNTAwLAogICAgICAic2lnbmVyIjogewogICAgICAgICJwdWJsaWNLZXkiOiAiMHhiMGQzNWNjZjcwYTZkODRlMmJjYTFkYzE2NmE0YzMzMjRkN2VkZDg2ZTg3OWFkZDJiYTY1MTFjOGVmNmJmZDg5YTE1NTM0ZTY3NDY3Y2NkOWM5MjExNTM0YjMzMjk1YTEiLAogICAgICAgICJwcm9vZk9mUG9zc2Vzc2lvbiI6ICIweGE0YjQ4MGE5YTA3YjRhOTc2MzBkZDlkNmUyYmY0ODM0YTNlNjcwNDQ4YjU1NzVlM2JhNzJhMDNlMDZlYzUwOWVjODU5ODQwYTExMDRiMThmMGNkNTQ2OTZlNmY5OGFkYjBlOWY1MjYwMTYxMzMyZmUzMmE1MGNiMWE5ODA2YjFiNTAyNTAzNzczMWVhNzdjNjQxZDYwN2ZkMDU4NGNlMjlkNzk1NGY1ZThmNTIzYzEzYTJlNTczMjUxMTIyN2Y1MCIKICAgICAgfQogICAgfSwKICAgIHsKICAgICAgIm5vZGVJRCI6ICJOb2RlSUQtN2VSdm5mczJhMlB2clBIVXVDUlJwUFZBb1ZqYld4YUZHIiwKICAgICAgInJld2FyZEFkZHJlc3MiOiAiWC1jdXN0b20xdjZ2dXd4amdyMDQzc2cwbnV1aHE3MGs2dmdudWxlNjkyZnZucjkiLAogICAgICAiZGVsZWdhdGlvbkZlZSI6IDYyNTAwLAogICAgICAic2lnbmVyIjogewogICAgICAgICJwdWJsaWNLZXkiOiAiMHhhNWZkZTYwNDJjNmUwZWU0ODJmNDYzNDZkZjA0NjAwMGNkNTdkZDg3OGQzNjYzN2E1YTYyYWRlYzA3YTUxMTRjZGVlYTA5NGE4NWY0ZjcyYjQ2NjQ1Zjk0ZTkwNzY2OTIiLAogICAgICAgICJwcm9vZk9mUG9zc2Vzc2lvbiI6ICIweDkxNDQ0MjMwYzVjZWI4ZTUxNjQyMTM5ZTE4NDJiNTZmODU2Mzg2NTM3NmI2ZjQyMDViZWNhNGRjMGJiMGJjNGIzMTRiY2UxZTE5ZTNiNTQyYTM5NDFlY2U1MWFlMjA1ZTAzYTA5NDgyNGZlZTI4ZjlmNzAyZWQzMTA3NTZmMDYzN2JmMTY2MzcxNjU2ZTFjM2ViOTAwMWRmODlmNGNkY2NjNzM0MTAyNDJhNmQ4NzVlYjYzNjNkMTJiY2U0MDMxNiIKICAgICAgfQogICAgfQogIF0sCiAgImNDaGFpbkdlbmVzaXMiOiAie1wiY29uZmlnXCI6e1wiY2hhaW5JZFwiOjQzMTE3LFwiaG9tZXN0ZWFkQmxvY2tcIjowLFwiZGFvRm9ya0Jsb2NrXCI6MCxcImRhb0ZvcmtTdXBwb3J0XCI6dHJ1ZSxcImVpcDE1MEJsb2NrXCI6MCxcImVpcDE1MEhhc2hcIjpcIjB4MjA4Njc5OWFlZWJlYWUxMzVjMjQ2YzY1MDIxYzgyYjRlMTVhMmM0NTEzNDA5OTNhYWNmZDI3NTE4ODY1MTRmMFwiLFwiZWlwMTU1QmxvY2tcIjowLFwiZWlwMTU4QmxvY2tcIjowLFwiYnl6YW50aXVtQmxvY2tcIjowLFwiY29uc3RhbnRpbm9wbGVCbG9ja1wiOjAsXCJwZXRlcnNidXJnQmxvY2tcIjowLFwiaXN0YW5idWxCbG9ja1wiOjAsXCJtdWlyR2xhY2llckJsb2NrXCI6MH0sXCJub25jZVwiOlwiMHgwXCIsXCJ0aW1lc3RhbXBcIjpcIjB4MFwiLFwiZXh0cmFEYXRhXCI6XCIweDAwXCIsXCJnYXNMaW1pdFwiOlwiMHg1ZjVlMTAwXCIsXCJkaWZmaWN1bHR5XCI6XCIweDBcIixcIm1peEhhc2hcIjpcIjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMFwiLFwiY29pbmJhc2VcIjpcIjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMFwiLFwiYWxsb2NcIjp7XCIwMTAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwXCI6e1wiY29kZVwiOlwiMHg3MzAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAzMDE0NjA4MDYwNDA1MjYwMDQzNjEwNjAzZDU3NjAwMDM1NjBlMDFjODA2MzFlMDEwNDM5MTQ2MDQyNTc4MDYzYjY1MTBiYjMxNDYwNmU1NzViNjAwMDgwZmQ1YjYwNWM2MDA0ODAzNjAzNjAyMDgxMTAxNTYwNTY1NzYwMDA4MGZkNWI1MDM1NjBiMTU2NWI2MDQwODA1MTkxODI1MjUxOTA4MTkwMDM2MDIwMDE5MGYzNWI4MTgwMTU2MDc5NTc2MDAwODBmZDViNTA2MGFmNjAwNDgwMzYwMzYwODA4MTEwMTU2MDhlNTc2MDAwODBmZDViNTA2MDAxNjAwMTYwYTAxYjAzODEzNTE2OTA2MDIwODEwMTM1OTA2MDQwODEwMTM1OTA2MDYwMDEzNTYwYjY1NjViMDA1YjMwY2Q5MDU2NWI4MzYwMDE2MDAxNjBhMDFiMDMxNjgxODM2MTA4ZmM4NjkwODExNTAyOTA2MDQwNTE2MDAwNjA0MDUxODA4MzAzODE4ODg4ODc4YzhhY2Y5NTUwNTA1MDUwNTA1MDE1ODAxNTYwZjQ1NzNkNjAwMDgwM2UzZDYwMDBmZDViNTA1MDUwNTA1MDU2ZmVhMjY0Njk3MDY2NzM1ODIyMTIyMDFlZWJjZTk3MGZlM2Y1Y2I5NmJmOGFjNmJhNWY1YzEzM2ZjMjkwOGFlM2RjZDUxMDgyY2ZlZThmNTgzNDI5ZDA2NDczNmY2YzYzNDMwMDA2MGEwMDMzXCIsXCJiYWxhbmNlXCI6XCIweDBcIn0sXCIweDY0M0YyNDU0NDMwRTIxODc1MGI1ZTY1MzNkOUMwZTBEZDUwQjhkNjhcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn0sXCIweGY5QkZBNEM0NWE4ZDgzMGE1OTFCMzM3NDMyMGZkOENDRjNGRDc1RDRcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn0sXCIweEQ5ZDRmMTZhNzFFMjNlRGY4ZTJGMmExRWJlY2Q0NkIwMzE3N2EyMmNcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn0sXCIweDJhMTc4MzE0MjViYzZEMjAwODREMTUyNmIxMDAxQzQ1MUVENEM0QTdcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn0sXCIweDdjNUE4NjM5RjFlODZGMTM0ZjFFNDIzOTQyOWY3NTZBMTQ0MWUzMjJcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn0sXCIweGZEREVmNWNiMEQwOUU0ODNkQkFCNTg3QkE5NTg2NTdCNzlBNDJFNThcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn0sXCIweEI0Y0E2QzEyMUQ2Mjg3YWY3YWM3Y2I2MkFlMzNkMmIwNTRiOUZDNDRcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn0sXCIweEM3MUE2MWE4MTVlNDlkMTZDNDI1NDgyQTM0MmEzNjdDRDQyRTM4YTZcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn19LFwibnVtYmVyXCI6XCIweDBcIixcImdhc1VzZWRcIjpcIjB4MFwiLFwicGFyZW50SGFzaFwiOlwiMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwXCJ9IiwKICAibWVzc2FnZSI6ICJFdG5hIGhlcmUgd2UgY29tZSIKfQo="
```

The out of this command includes your Node's `NodeID`, `BLS PublicKey`, and `BLS PoP` as the first thing under the large "Avalanche" ascii art. It should look something like:

```zsh
     _____               .__                       .__
    /  _  \___  _______  |  | _____    ____   ____ |  |__   ____    ,_ o
   /  /_\  \  \/ /\__  \ |  | \__  \  /    \_/ ___\|  |  \_/ __ \   / //\,
  /    |    \   /  / __ \|  |__/ __ \|   |  \  \___|   Y  \  ___/    \>> |
  \____|__  /\_/  (____  /____(____  /___|  /\___  >___|  /\___  >    \\
          \/           \/          \/     \/     \/     \/     \/
[10-10|17:31:24.809] INFO node/node.go:143 initializing node {"version": "avalanchego/1.11.11", "nodeID": "NodeID-GR7FiVZaXjbB78d4GFPCawq1nJwA3Lr2d", "stakingKeyType": "RSA", "nodePOP": {"publicKey":"0xa2d5200c9a468e0acf4dd8e2330cd736589767d613848683f7f70527e18e374f8f227b364024af79775dc371e989898","proofOfPossession":"0xb917e55a6d97bc77b9921662724d4ac14002268ded9d50f5c209bb54fc28064027a05ceef6e163a675a5bc1780571d2913d77fb495024311d07cf144fb91237be4d91cb36b6bdbe2967c0734d22217e3acf5ad26bfa58e2bc8305815e89898989"}, ...
```

## Build the CLI

In a separate terminal window:

```zsh
git clone https://github.com/ava-labs/avalanche-cli.git
cd avalanche-cli
git checkout acp-77
./scripts/build.sh
```

Next:

`./bin/avalanche blockchain create <chainName>`

Choose your configs, and for ease of use, just use the `ewoq` key for everything.

Then:

`./bin/avalanche blockchain deploy <chainName> --devnet --endpoint https://etna.avax-dev.network`

```zsh
? How many bootstrap validators do you want to set up?: 1
? Have you set up your own Avalanche Nodes?:
  ▸ Yes
✗ What is the NodeID of the node you want to add as bootstrap validator?: <NodeID>
✗ What is the node's BLS public key?: <BLS Public Key>
✔ What is the node's BLS proof of possession?: <BLS PoP>
✔ Get address from an existing stored key (created from avalanche key create or avalanche key import)
✔ ewoq
```

Ignore any error messages about nodes, and **note** the deployment results:

```zsh
+--------------------+---------------------------------------------------+
| DEPLOYMENT RESULTS |                                                   |
+--------------------+---------------------------------------------------+
| Chain Name         | test                                              |
+--------------------+---------------------------------------------------+
| Subnet ID          | kawEpA4hUhfH6Ke39zJDYkt82AiLdDZueKurvxQT9zQAPc456 |
+--------------------+---------------------------------------------------+
| VM ID              | qBNid6dean7PHGiAicRFmAoS3VVzi8iV9NRjU5PhQdX8yL456 |
+--------------------+---------------------------------------------------+
| Blockchain ID      | iii8L591ceGgER7nk7ZzpyjopRx3XvShMUkVjvSP8o514t456 |
+--------------------+                                                   +
| P-Chain TXID       |                                                   |
+--------------------+---------------------------------------------------+
```

## Build the Subnet-EVM Binary

In a third terminal window, ensuring you have replaced `<VM ID>` with the VM ID you got from the previous step:

```zsh
cd $GOPATH/src/github.com/ava-labs
git clone https://github.com/ava-labs/subnet-evm.git
cd $GOPATH/src/github.com/ava-labs/subnet-evm
./scripts/build.sh ~/.avalanchego/plugins/<VM ID>
```

## Stop your Node

In the terminal window running your avalanchego node, stop it by pressing `Ctrl + C`

### Restart your Node

Restart the Node to track your L1, tracking the subnet with the flag `--track-subnet="<SubnetID>"` appended to the end. Replace `<Subnet ID>` with the value from above.

```zsh
➜ avalanchego git:(v1.12.0-initial-poc.5) ./build/avalanchego \
    --network-id="network-76" \
    --bootstrap-ids="NodeID-8LbTmmGsDC991SbD8Nkx88VULT3XYzYXC,NodeID-bojBKDrpt81bYhxYKQfLw89V7CpoH2m7,NodeID-WrLWMK5sJ4dBUAsx1dP2FUyTqrYwbFA1,NodeID-DDhXtFm6Q9tCq2yiFRmcSMKvHgUgh8yQC,NodeID-QDYnWDQd6g4cQ5H6yiWNqSmfRMBqEH9AG" \
    --bootstrap-ips="52.201.126.172:9651,34.233.248.130:9651,107.21.11.213:9651,35.170.144.5:9651,98.82.41.186:9651" \
    --upgrade-file-content="ewogICAgImFwcmljb3RQaGFzZTFUaW1lIjogIjIwMjAtMTItMDVUMDU6MDA6MDBaIiwKICAgICJhcHJpY290UGhhc2UyVGltZSI6ICIyMDIwLTEyLTA1VDA1OjAwOjAwWiIsCiAgICAiYXByaWNvdFBoYXNlM1RpbWUiOiAiMjAyMC0xMi0wNVQwNTowMDowMFoiLAogICAgImFwcmljb3RQaGFzZTRUaW1lIjogIjIwMjAtMTItMDVUMDU6MDA6MDBaIiwKICAgICJhcHJpY290UGhhc2U0TWluUENoYWluSGVpZ2h0IjogMCwKICAgICJhcHJpY290UGhhc2U1VGltZSI6ICIyMDIwLTEyLTA1VDA1OjAwOjAwWiIsCiAgICAiYXByaWNvdFBoYXNlUHJlNlRpbWUiOiAiMjAyMC0xMi0wNVQwNTowMDowMFoiLAogICAgImFwcmljb3RQaGFzZTZUaW1lIjogIjIwMjAtMTItMDVUMDU6MDA6MDBaIiwKICAgICJhcHJpY290UGhhc2VQb3N0NlRpbWUiOiAiMjAyMC0xMi0wNVQwNTowMDowMFoiLAogICAgImJhbmZmVGltZSI6ICIyMDIwLTEyLTA1VDA1OjAwOjAwWiIsCiAgICAiY29ydGluYVRpbWUiOiAiMjAyMC0xMi0wNVQwNTowMDowMFoiLAogICAgImNvcnRpbmFYQ2hhaW5TdG9wVmVydGV4SUQiOiAiMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTFMcG9ZWSIsCiAgICAiZHVyYW5nb1RpbWUiOiAiMjAyMC0xMi0wNVQwNTowMDowMFoiLAogICAgImV0bmFUaW1lIjogIjIwMjQtMTAtMDlUMjA6MDA6MDBaIgp9Cg==" \
    --genesis-file-content="ewogICJuZXR3b3JrSUQiOiA3NiwKICAiYWxsb2NhdGlvbnMiOiBbCiAgICB7CiAgICAgICJldGhBZGRyIjogIjB4QzcxQTYxYTgxNWU0OWQxNkM0MjU0ODJBMzQyYTM2N0NENDJFMzhhNiIsCiAgICAgICJhdmF4QWRkciI6ICJYLWN1c3RvbTF2NnZ1d3hqZ3IwNDNzZzBudXVocTcwazZ2Z251bGU2OTJmdm5yOSIsCiAgICAgICJpbml0aWFsQW1vdW50IjogNTAwMDAwMDAwMDAwMDAwMDAwLAogICAgICAidW5sb2NrU2NoZWR1bGUiOiBbCiAgICAgICAgewogICAgICAgICAgImFtb3VudCI6IDEwMDAwMDAwMDAwMDAwMDAwMCwKICAgICAgICAgICJsb2NrdGltZSI6IDE2MzM4MjQwMDAKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJhbW91bnQiOiAxMDAwMDAwMDAwMDAwMDAwMDAsCiAgICAgICAgICAibG9ja3RpbWUiOiAxNjMzODI1MDAwCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAiYW1vdW50IjogMTAwMDAwMDAwMDAwMDAwMDAwLAogICAgICAgICAgImxvY2t0aW1lIjogMTYzMzgyNjAwMAogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgImFtb3VudCI6IDEwMDAwMDAwMDAwMDAwMDAwMCwKICAgICAgICAgICJsb2NrdGltZSI6IDE2MzM4MjcwMDAKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJhbW91bnQiOiAxMDAwMDAwMDAwMDAwMDAwMDAsCiAgICAgICAgICAibG9ja3RpbWUiOiAxNjMzODI4MDAwCiAgICAgICAgfQogICAgICBdCiAgICB9CiAgXSwKICAic3RhcnRUaW1lIjogMTcyNTMwMDAwMCwKICAiaW5pdGlhbFN0YWtlRHVyYXRpb24iOiAzMTUzMDAwMCwKICAiaW5pdGlhbFN0YWtlRHVyYXRpb25PZmZzZXQiOiA1NDAwLAogICJpbml0aWFsU3Rha2VkRnVuZHMiOiBbCiAgICAiWC1jdXN0b20xdjZ2dXd4amdyMDQzc2cwbnV1aHE3MGs2dmdudWxlNjkyZnZucjkiCiAgXSwKICAiaW5pdGlhbFN0YWtlcnMiOiBbCiAgICB7CiAgICAgICJub2RlSUQiOiAiTm9kZUlELWdwWFdCRXhRU1pYcUpQUXQ2TDZNbnZlVWZncjdISjRxIiwKICAgICAgInJld2FyZEFkZHJlc3MiOiAiWC1jdXN0b20xdjZ2dXd4amdyMDQzc2cwbnV1aHE3MGs2dmdudWxlNjkyZnZucjkiLAogICAgICAiZGVsZWdhdGlvbkZlZSI6IDYyNTAwLAogICAgICAic2lnbmVyIjogewogICAgICAgICJwdWJsaWNLZXkiOiAiMHhhMTRkNjdmMDk3ZDdlNjUxNDY5NmZkODMwODA3OTRiNmI1Y2E2NjQwMDFmMmVkZTRmZDZmMDFkYTQ5MzNkYjg3NWZmMDI4ZmVjNDJiMjlmYzU1MjQ5NDFlMGYyMDgzMGYiLAogICAgICAgICJwcm9vZk9mUG9zc2Vzc2lvbiI6ICIweDgyMzUyYWUxZTAxMDM4MTczZTkyZTA4OGJkMzRjMmJlZTljYzRiMzRkZjVjNWU4YmQyNzczY2VmOTIzOGVlZjg3MGMyZjkzZmE4OTYwNzMzMmNjYmI4NGFhNjY2MDhjNzA2YjdjMmYxMjdiOGI4MGM0NjFjMDRiYmM2MDgyYWZiZmZlMjIwYWFjNzlmNjY1MzNlYTdjNjNmMDQ1MWQ3ZDMyNDU2MzY5ZGQzMzVjOTcxMDkzOGVlNDExMWQwOGQ3OSIKICAgICAgfQogICAgfSwKICAgIHsKICAgICAgIm5vZGVJRCI6ICJOb2RlSUQtNzhpYldwanRaejVaR1Q2RXlURWR1OFZLbWJvVUhUdUdUIiwKICAgICAgInJld2FyZEFkZHJlc3MiOiAiWC1jdXN0b20xdjZ2dXd4amdyMDQzc2cwbnV1aHE3MGs2dmdudWxlNjkyZnZucjkiLAogICAgICAiZGVsZWdhdGlvbkZlZSI6IDYyNTAwLAogICAgICAic2lnbmVyIjogewogICAgICAgICJwdWJsaWNLZXkiOiAiMHg4MzI3ZGJlMWJhNDExYzI3MDYzN2IwODBhODQ3MWZiNDFlZWI4YTliMzkxN2FmMDcyNzUwMWVmOGJkYWE5MDFkMDYzNzgwYmQ3MDJmMzBmNDU4YTYxZjNkNDI5N2RjOTgiLAogICAgICAgICJwcm9vZk9mUG9zc2Vzc2lvbiI6ICIweGE5YzAzOWI1NzY1YWIwNjhiZDYzMmJiY2RjOWJjMmE1M2YyOWUyYzU2YjMzZTMwZDczMmEyM2Q4YzQzMGQ1M2VmNDdlYmNjZmFhNWNmY2VkZDhmMDQxYzJjMTM0OGYwYjBlYWM0MTMxOTJiNzU0NGQyODRmODJkMWZhMGY3NGY5OGQ1ODA1OTA1MzYzYjgxODZlZmRlZjZlNzcxODJmYjFlNzE0N2Y4NTExZTkwMGQxOTVkYjA2ZGE2YTIyZjBhMCIKICAgICAgfQogICAgfSwKICAgIHsKICAgICAgIm5vZGVJRCI6ICJOb2RlSUQtTDRDWThCNXVWU0RlNGNuTjFCcGVEc0hhY01wNHE0cThxIiwKICAgICAgInJld2FyZEFkZHJlc3MiOiAiWC1jdXN0b20xdjZ2dXd4amdyMDQzc2cwbnV1aHE3MGs2dmdudWxlNjkyZnZucjkiLAogICAgICAiZGVsZWdhdGlvbkZlZSI6IDYyNTAwLAogICAgICAic2lnbmVyIjogewogICAgICAgICJwdWJsaWNLZXkiOiAiMHhhOThjNjQ2YThjODYyZWMxNTMyNmU0Y2ZlMmEwZjY2YThmYjdjZjU1NTc2NWY4M2ZmMzIwYTFhNzYyNjgyMjhmM2M4YjI2MmQxZGU0MDA4ZTBiYTQ5YTg5Y2ZhYmZiOTUiLAogICAgICAgICJwcm9vZk9mUG9zc2Vzc2lvbiI6ICIweDk1YzgxNmE0ZDI5MmE0N2M0ZDk5MzRlMzU4NjIxZDA3ODVmZDI5MjBhMWEzMDRjZWJjOWI3ZTQ3NDE3ZTNmZmY3ODBjZmNkZGY3Y2ExMTc3YjQ1YmJmYWMyZjk5Nzg1ODE3NjFkOWRkZDU1ZWM2MTQyZDkyOTk4ZWVhZGJhZmU4Y2Q3NjUxMDU2ZmJiNzlhZmVhNjQzZjBjZDIwZmY0ZjYzODlkZGQ5MWVlMmRiNDU3OTQzOGE2OTA4NjA5YjRjMSIKICAgICAgfQogICAgfSwKICAgIHsKICAgICAgIm5vZGVJRCI6ICJOb2RlSUQtUDVRR0g0RVhkZHJjeU5BemtxeVpLSFhnRXBWWDZIRXhMIiwKICAgICAgInJld2FyZEFkZHJlc3MiOiAiWC1jdXN0b20xdjZ2dXd4amdyMDQzc2cwbnV1aHE3MGs2dmdudWxlNjkyZnZucjkiLAogICAgICAiZGVsZWdhdGlvbkZlZSI6IDYyNTAwLAogICAgICAic2lnbmVyIjogewogICAgICAgICJwdWJsaWNLZXkiOiAiMHhiMGQzNWNjZjcwYTZkODRlMmJjYTFkYzE2NmE0YzMzMjRkN2VkZDg2ZTg3OWFkZDJiYTY1MTFjOGVmNmJmZDg5YTE1NTM0ZTY3NDY3Y2NkOWM5MjExNTM0YjMzMjk1YTEiLAogICAgICAgICJwcm9vZk9mUG9zc2Vzc2lvbiI6ICIweGE0YjQ4MGE5YTA3YjRhOTc2MzBkZDlkNmUyYmY0ODM0YTNlNjcwNDQ4YjU1NzVlM2JhNzJhMDNlMDZlYzUwOWVjODU5ODQwYTExMDRiMThmMGNkNTQ2OTZlNmY5OGFkYjBlOWY1MjYwMTYxMzMyZmUzMmE1MGNiMWE5ODA2YjFiNTAyNTAzNzczMWVhNzdjNjQxZDYwN2ZkMDU4NGNlMjlkNzk1NGY1ZThmNTIzYzEzYTJlNTczMjUxMTIyN2Y1MCIKICAgICAgfQogICAgfSwKICAgIHsKICAgICAgIm5vZGVJRCI6ICJOb2RlSUQtN2VSdm5mczJhMlB2clBIVXVDUlJwUFZBb1ZqYld4YUZHIiwKICAgICAgInJld2FyZEFkZHJlc3MiOiAiWC1jdXN0b20xdjZ2dXd4amdyMDQzc2cwbnV1aHE3MGs2dmdudWxlNjkyZnZucjkiLAogICAgICAiZGVsZWdhdGlvbkZlZSI6IDYyNTAwLAogICAgICAic2lnbmVyIjogewogICAgICAgICJwdWJsaWNLZXkiOiAiMHhhNWZkZTYwNDJjNmUwZWU0ODJmNDYzNDZkZjA0NjAwMGNkNTdkZDg3OGQzNjYzN2E1YTYyYWRlYzA3YTUxMTRjZGVlYTA5NGE4NWY0ZjcyYjQ2NjQ1Zjk0ZTkwNzY2OTIiLAogICAgICAgICJwcm9vZk9mUG9zc2Vzc2lvbiI6ICIweDkxNDQ0MjMwYzVjZWI4ZTUxNjQyMTM5ZTE4NDJiNTZmODU2Mzg2NTM3NmI2ZjQyMDViZWNhNGRjMGJiMGJjNGIzMTRiY2UxZTE5ZTNiNTQyYTM5NDFlY2U1MWFlMjA1ZTAzYTA5NDgyNGZlZTI4ZjlmNzAyZWQzMTA3NTZmMDYzN2JmMTY2MzcxNjU2ZTFjM2ViOTAwMWRmODlmNGNkY2NjNzM0MTAyNDJhNmQ4NzVlYjYzNjNkMTJiY2U0MDMxNiIKICAgICAgfQogICAgfQogIF0sCiAgImNDaGFpbkdlbmVzaXMiOiAie1wiY29uZmlnXCI6e1wiY2hhaW5JZFwiOjQzMTE3LFwiaG9tZXN0ZWFkQmxvY2tcIjowLFwiZGFvRm9ya0Jsb2NrXCI6MCxcImRhb0ZvcmtTdXBwb3J0XCI6dHJ1ZSxcImVpcDE1MEJsb2NrXCI6MCxcImVpcDE1MEhhc2hcIjpcIjB4MjA4Njc5OWFlZWJlYWUxMzVjMjQ2YzY1MDIxYzgyYjRlMTVhMmM0NTEzNDA5OTNhYWNmZDI3NTE4ODY1MTRmMFwiLFwiZWlwMTU1QmxvY2tcIjowLFwiZWlwMTU4QmxvY2tcIjowLFwiYnl6YW50aXVtQmxvY2tcIjowLFwiY29uc3RhbnRpbm9wbGVCbG9ja1wiOjAsXCJwZXRlcnNidXJnQmxvY2tcIjowLFwiaXN0YW5idWxCbG9ja1wiOjAsXCJtdWlyR2xhY2llckJsb2NrXCI6MH0sXCJub25jZVwiOlwiMHgwXCIsXCJ0aW1lc3RhbXBcIjpcIjB4MFwiLFwiZXh0cmFEYXRhXCI6XCIweDAwXCIsXCJnYXNMaW1pdFwiOlwiMHg1ZjVlMTAwXCIsXCJkaWZmaWN1bHR5XCI6XCIweDBcIixcIm1peEhhc2hcIjpcIjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMFwiLFwiY29pbmJhc2VcIjpcIjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMFwiLFwiYWxsb2NcIjp7XCIwMTAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwXCI6e1wiY29kZVwiOlwiMHg3MzAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAzMDE0NjA4MDYwNDA1MjYwMDQzNjEwNjAzZDU3NjAwMDM1NjBlMDFjODA2MzFlMDEwNDM5MTQ2MDQyNTc4MDYzYjY1MTBiYjMxNDYwNmU1NzViNjAwMDgwZmQ1YjYwNWM2MDA0ODAzNjAzNjAyMDgxMTAxNTYwNTY1NzYwMDA4MGZkNWI1MDM1NjBiMTU2NWI2MDQwODA1MTkxODI1MjUxOTA4MTkwMDM2MDIwMDE5MGYzNWI4MTgwMTU2MDc5NTc2MDAwODBmZDViNTA2MGFmNjAwNDgwMzYwMzYwODA4MTEwMTU2MDhlNTc2MDAwODBmZDViNTA2MDAxNjAwMTYwYTAxYjAzODEzNTE2OTA2MDIwODEwMTM1OTA2MDQwODEwMTM1OTA2MDYwMDEzNTYwYjY1NjViMDA1YjMwY2Q5MDU2NWI4MzYwMDE2MDAxNjBhMDFiMDMxNjgxODM2MTA4ZmM4NjkwODExNTAyOTA2MDQwNTE2MDAwNjA0MDUxODA4MzAzODE4ODg4ODc4YzhhY2Y5NTUwNTA1MDUwNTA1MDE1ODAxNTYwZjQ1NzNkNjAwMDgwM2UzZDYwMDBmZDViNTA1MDUwNTA1MDU2ZmVhMjY0Njk3MDY2NzM1ODIyMTIyMDFlZWJjZTk3MGZlM2Y1Y2I5NmJmOGFjNmJhNWY1YzEzM2ZjMjkwOGFlM2RjZDUxMDgyY2ZlZThmNTgzNDI5ZDA2NDczNmY2YzYzNDMwMDA2MGEwMDMzXCIsXCJiYWxhbmNlXCI6XCIweDBcIn0sXCIweDY0M0YyNDU0NDMwRTIxODc1MGI1ZTY1MzNkOUMwZTBEZDUwQjhkNjhcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn0sXCIweGY5QkZBNEM0NWE4ZDgzMGE1OTFCMzM3NDMyMGZkOENDRjNGRDc1RDRcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn0sXCIweEQ5ZDRmMTZhNzFFMjNlRGY4ZTJGMmExRWJlY2Q0NkIwMzE3N2EyMmNcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn0sXCIweDJhMTc4MzE0MjViYzZEMjAwODREMTUyNmIxMDAxQzQ1MUVENEM0QTdcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn0sXCIweDdjNUE4NjM5RjFlODZGMTM0ZjFFNDIzOTQyOWY3NTZBMTQ0MWUzMjJcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn0sXCIweGZEREVmNWNiMEQwOUU0ODNkQkFCNTg3QkE5NTg2NTdCNzlBNDJFNThcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn0sXCIweEI0Y0E2QzEyMUQ2Mjg3YWY3YWM3Y2I2MkFlMzNkMmIwNTRiOUZDNDRcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn0sXCIweEM3MUE2MWE4MTVlNDlkMTZDNDI1NDgyQTM0MmEzNjdDRDQyRTM4YTZcIjp7XCJiYWxhbmNlXCI6XCIweDE0MzFFMEZBRTZENzIxN0NBQTAwMDAwMDBcIn19LFwibnVtYmVyXCI6XCIweDBcIixcImdhc1VzZWRcIjpcIjB4MFwiLFwicGFyZW50SGFzaFwiOlwiMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwXCJ9IiwKICAibWVzc2FnZSI6ICJFdG5hIGhlcmUgd2UgY29tZSIKfQo=" \
--track-subnets="<Subnet ID>"
```

So long as you don't see any errors in your logs, you are now running a L1 Validator locally on your machine.

## Confirmation

In a fourth terminal:

```zsh
curl -X POST --data '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"eth_getChainConfig",
    "params" :[]
}' -H 'content-type:application/json;' 127.0.0.1:9650/ext/bc/<BlockchainID>/rpc
```