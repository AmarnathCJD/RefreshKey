```markdown
# Hotkey Refresh Rate Changer

This script allows you to toggle between available refresh rates using a hotkey. It uses the `wmic` command to retrieve the current refresh rate of the display and then switches between the available refresh rates based on the current value.

## Setting Hotkey Using Regedit

1. Press `Win + R` to open the Run dialog, type `regedit`, and press Enter to open the Registry Editor.

2. Navigate to `HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\AppKey\18`.

3. Right-click on the right pane, select `New > String Value`, and name the new string value `ShellExecute`.

4. Double-click on the `ShellExecute` value and enter the full path to the `runcsr.vbs` script.

5. Close the Registry Editor.

After making these changes, the `runcsr.vbs` script should be executed when the specified hotkey (AppKey 18) is pressed.
```

## Script Details

The provided batch script uses the `wmic` command to get the current refresh rate of the display. It then checks the current refresh rate and uses the `csr` command to switch between the available refresh rates based on the current value.

```batch
@echo off
for /f "tokens=2 delims==" %%I in ('wmic PATH Win32_videocontroller get currentrefreshrate /value') do set "refresh=%%I"
if %refresh%==144 (
    csr /f=60 /d=0
) else if %refresh%==60 (
    csr /f=144 /d=0
)
```

### (c) RoseLoverX

## License

This script is provided under the [MIT License](https://opensource.org/licenses/MIT).