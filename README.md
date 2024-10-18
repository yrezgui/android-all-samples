# android-all-samples

## Instructions

### Import repositories as submodules
```bash
git submodule add git@github.com:android/ai-samples.git samples/ai-samples
git submodule add git@github.com:android/animation-samples.git samples/animation-samples
git submodule add git@github.com:android/app-bundle-samples.git samples/app-bundle-samples
git submodule add git@github.com:android/architecture-components-samples.git samples/architecture-components-samples
git submodule add git@github.com:android/camera-samples.git samples/camera-samples
git submodule add git@github.com:android/car-samples.git samples/car-samples
git submodule add git@github.com:android/compose-samples.git samples/compose-samples
git submodule add git@github.com:android/connectivity-samples.git samples/connectivity-samples
git submodule add git@github.com:android/databinding-samples.git samples/databinding-samples
git submodule add git@github.com:android/fit-samples.git samples/fit-samples
git submodule add git@github.com:android/graphics-samples.git samples/graphics-samples
git submodule add git@github.com:android/health-samples.git samples/health-samples
git submodule add git@github.com:android/identity-samples.git samples/identity-samples
git submodule add git@github.com:android/input-samples.git samples/input-samples
git submodule add git@github.com:android/kotlin-multiplatform-samples.git samples/kotlin-multiplatform-samples
git submodule add git@github.com:android/media-samples.git samples/media-samples
git submodule add git@github.com:android/midi-samples.git samples/midi-samples
git submodule add git@github.com:android/neural-networks-samples.git samples/neural-networks-samples
git submodule add git@github.com:android/packager-manager-samples.git samples/packager-manager-samples
git submodule add git@github.com:android/performance-samples.git samples/performance-samples
git submodule add git@github.com:android/privacy-sandbox-samples.git samples/privacy-sandbox-samples
git submodule add git@github.com:android/renderscript-samples.git samples/renderscript-samples
git submodule add git@github.com:android/search-samples.git samples/search-samples
git submodule add git@github.com:android/security-samples.git samples/security-samples
git submodule add git@github.com:android/sensors-samples.git samples/sensors-samples
git submodule add git@github.com:android/storage-samples.git samples/storage-samples
git submodule add git@github.com:android/testing-samples.git samples/testing-samples
git submodule add git@github.com:android/tv-samples.git samples/tv-samples
git submodule add git@github.com:android/uamp.git samples/uamp
git submodule add git@github.com:android/user-interface-samples.git samples/user-interface-samples
git submodule add git@github.com:android/views-widgets-samples.git samples/views-widgets-samples
git submodule add git@github.com:android/wear-os-samples.git samples/wear-os-samples
# Use branch system
# git submodule add git@github.com:android/architecture-samples.git samples/architecture-samples
# Massive catalog system
# git submodule add git@github.com:android/platform-samples.git samples/platform-samples
# Special build system
# git submodule add git@github.com:android/ndk-samples.git samples/ndk-samples
# git submodule add git@github.com:android/games-samples.git samples/games-samples
# Archived
# git submodule add git@github.com:android/permissions-samples.git samples/permissions-samples-samples
# git submodule add git@github.com:android/location-samples.git samples/location-samples
# git submodule add git@github.com:android/enterprise-samples.git samples/enterprise-samples
# git submodule add git@github.com:android/app-actions-samples.git samples/app-actions-samples
```

### Count all file extensions
```bash
find . -type f | sed -n 's/.*\.//p' | sort | uniq -c | awk '{print $2 ": " $1}'
```

## Get last git commit timestamp for current folder
```bash
git log -1 --format=%ct
```