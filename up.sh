git pull
cp ../frida_shared/cmake-build-release/libfrida_shared.dylib ./libfrida_darwin_amd64.dylib
git add .
git commit -m "macos"
git push -u origin master
