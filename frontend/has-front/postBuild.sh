# Clear old files
rm -dr ../../backend/server/ui/dist

# Recreate the folder
mkdir ../../backend/server/ui/dist

# Copy dist files
cp -r dist/* ../../backend/server/ui/dist/