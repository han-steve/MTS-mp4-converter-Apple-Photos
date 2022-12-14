# MTS-to-MP4 Converter

### Converts your old MTS videos to a Apple Quick Time compatible mp4 format and copies over the metadata (such as dates). Perfect for importing old videos to Apple Photos!

My parents left me a hard drive full of my childhood videos. Instead of leaving them in the folder, I wanted to import them into Apple Photos so I could have all my photos in one place. However, the videos are in MTS format, which Apple Photos doesn't support. I tried multiple ways of converting them, but getting the timestamp and audio correct as well as getting the converted file to be compatible with QuickTime was difficult. I finally found the following commands that work well for me:

```
ffmpeg -i NAME.MTS -c:v libx265 -preset fast -crf 28 -tag:v hvc1 -c:a eac3 -b:a 224k output.mp4

exiftool -tagsFromFile "NAME.MTS" -time:all "output.mp4"
```

The first command converts the video and autio to a QuickTime-compatible format while the second one copies over the timestamp.  
To automate the process of converting all photos and preserving the folder structure, I wrote this Go script.

Make sure that you have ffmpeg and exiftool installed. If you are on a Mac, I recommend using homebrew for this. 
It's recommended to build it from source (by running go build cmd/main.go - if you don't have go installed, [install it](https://go.dev/doc/install)). 
To execute this command, provide the root directory, source folder name, and destination folder name as command line arguments. For example, execute ./main /Users/user/Desktop /Source /Destination.

After executing it, it shows a progress bar displaying the conversion progress. Since running ffmpeg on big files is slow, allow it plenty of time to run (for my library of 900 videos, it took 15 hours). 
