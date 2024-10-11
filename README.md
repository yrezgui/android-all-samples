# android-all-samples

## Instructions

### Import repositories as submodules
```bash
git submodule add git@github.com:android/ai-samples.git
git submodule add git@github.com:android/animation-samples.git
# Archived
# git submodule add git@github.com:android/app-actions-samples.git
git submodule add git@github.com:android/app-bundle-samples.git
git submodule add git@github.com:android/architecture-components-samples.git
# Use branch system
# git submodule add git@github.com:android/architecture-samples.git
git submodule add git@github.com:android/camera-samples.git
git submodule add git@github.com:android/car-samples.git
git submodule add git@github.com:android/compose-samples.git
git submodule add git@github.com:android/connectivity-samples.git
git submodule add git@github.com:android/databinding-samples.git
# Archived
# git submodule add git@github.com:android/enterprise-samples.git
git submodule add git@github.com:android/fit-samples.git
# Special build system
# git submodule add git@github.com:android/games-samples.git
git submodule add git@github.com:android/graphics-samples.git
git submodule add git@github.com:android/health-samples.git
git submodule add git@github.com:android/identity-samples.git
git submodule add git@github.com:android/input-samples.git
git submodule add git@github.com:android/kotlin-multiplatform-samples.git
# Archived
# git submodule add git@github.com:android/location-samples.git
git submodule add git@github.com:android/media-samples.git
git submodule add git@github.com:android/midi-samples.git
# Special build system
# git submodule add git@github.com:android/ndk-samples.git
git submodule add git@github.com:android/neural-networks-samples.git
git submodule add git@github.com:android/packager-manager-samples.git
git submodule add git@github.com:android/performance-samples.git
# Archived
# git submodule add git@github.com:android/permissions-samples.git
# Massive catalog system
# git submodule add git@github.com:android/platform-samples.git
git submodule add git@github.com:android/privacy-sandbox-samples.git
git submodule add git@github.com:android/renderscript-samples.git
git submodule add git@github.com:android/search-samples.git
git submodule add git@github.com:android/security-samples.git
git submodule add git@github.com:android/sensors-samples.git
git submodule add git@github.com:android/storage-samples.git
git submodule add git@github.com:android/testing-samples.git
git submodule add git@github.com:android/tv-samples.git
git submodule add git@github.com:android/uamp.git
git submodule add git@github.com:android/user-interface-samples.git
git submodule add git@github.com:android/views-widgets-samples.git
git submodule add git@github.com:android/wear-os-samples.git
```

### Pack all useful files
```bash
repopack --include "build.gradle, build.gradle.kts, local.properties, gradle.properties, samples/**/*.java,samples/**/*.kt, **/*.gradle, **/*.kts, **/AndroidManifest.xml" --ignore ".google, .gradle, gradle, .idea"
```

### Count all file extensions
```bash
find . -type f | sed -n 's/.*\.//p' | sort | uniq -c | awk '{print $2 ": " $1}'
```