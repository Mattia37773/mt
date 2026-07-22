#!/usr/bin/env bash

# ==============================================================================
# Directories 
# ==============================================================================

# Directories to monitor (Space-separated)
WATCH_DIRS=("cmd" "functions" "config")

# Files or extensions to ignore (Regex)
IGNORE_PATTERN="\.tmp$|\.log$"

# ==============================================================================
# Helpers
# ==============================================================================

# Green  
color_green() {
    tput setaf 46
    echo "$1"
    tput sgr0
}

# Dark Blue
color_darkblue() {
    tput setaf 21        
    echo "$1"
    tput sgr0
}

# Purple
color_purple() {
    tput setaf 165
    echo "$1"
    tput sgr0
}

# Light Blue 
color_lightblue() {
    tput setaf 51  
    echo "$1"
    tput sgr0
}

# Red 
color_red() {
    tput setaf 196
    echo "$1"
    tput sgr0
}


# ==============================================================================
# CONFIGURATION
# ==============================================================================

# The action to perform when a change is detected
# $1 = Path to the file that triggered the action
run_action() {
    local changed_file="$1"
    
    echo "-------------------------------------------------------"
    color_purple "[ACTION] Starting Build at: $(date +'%H:%M:%S')"
    color_darkblue "[INFO]   Triggered by: $changed_file"
    go build -ldflags "-X github.com/mattia37773/mt/config.Devmode=true"
    color_purple "[ACTION] Build Completed successfully."
    echo "-------------------------------------------------------"
    echo
}

# Polling interval in seconds
SLEEP_INTERVAL=1

# ==============================================================================
# INTERNAL LOGIC (Handle with care)
# ==============================================================================

if [[ "$(basename "$PWD")" == "bin" ]]; then
    cd ..
fi

declare -A FILE_SNAPSHOT

# Detects platform (Linux vs macOS) to use the correct stat command
stat_time() {
    if stat --version >/dev/null 2>&1; then
        stat -c "%Y" "$1" # GNU/Linux
    else
        stat -f "%m" "$1" # BSD/macOS
    fi
}

# Scans directories and populates the snapshot array
update_snapshot() {
    local -n snap=$1
    snap=() # Clear current snapshot
    for dir in "${WATCH_DIRS[@]}"; do
        if [ ! -d "$dir" ]; then 
            echo "[WARN] Directory not found: $dir"
            continue
        fi
        
        while IFS= read -r -d '' f; do
            # Skip files matching the ignore pattern
            if [[ -n "$IGNORE_PATTERN" && "$f" =~ $IGNORE_PATTERN ]]; then
                continue
            fi
            # Get last modified time
            snap["$f"]="$(stat_time "$f" 2>/dev/null || echo 0)"
        done < <(find "$dir" -type f -print0)
    done
}

echo "Watcher started. Monitoring: ${WATCH_DIRS[*]}"
echo "Press [CTRL+C] to stop."

# Initial baseline snapshot
update_snapshot FILE_SNAPSHOT

while true; do
    sleep "$SLEEP_INTERVAL"
    declare -A CURRENT
    update_snapshot CURRENT
    
    trigger_file=""

    # 1. Check for NEW or MODIFIED files
    for f in "${!CURRENT[@]}"; do
        if [[ -z "${FILE_SNAPSHOT[$f]:-}" ]]; then
            color_green "[CREATE] $f"
            trigger_file="$f"
            break
        elif [[ "${CURRENT[$f]}" != "${FILE_SNAPSHOT[$f]}" ]]; then
            color_lightblue "[MODIFY] $f"
            trigger_file="$f"
            break
        fi
    done

    # 2. Check for DELETED files (only if no trigger found yet)
    if [[ -z "$trigger_file" ]]; then
        for f in "${!FILE_SNAPSHOT[@]}"; do
            if [[ -z "${CURRENT[$f]:-}" ]]; then
                color_red "[DELETE] $f"
                trigger_file="$f"
                break
            fi
        done
    fi

    # Execute action if a change was detected
    if [[ -n "$trigger_file" ]]; then
        run_action "$trigger_file"
        
        # Refresh the master snapshot after the action 
        # This prevents build-artifacts from triggering a second loop
        update_snapshot FILE_SNAPSHOT
    fi
done
