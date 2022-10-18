# MTS-to-MP4 Converter

My parents left me a hard drive full of my childhood videos. Instead of leaving them in the folder, I wanted to import them into Apple Photos so I could have all my photos in one place. However, the videos are in MTS format, which Apple Photos doesn't support. I tried multiple ways of converting them, but getting the timestamp and audio correct as well as getting the converted file to be compatible with QuickTime was difficult. I finally found the following commands that work well for me:

```
 ffmpeg -i NAME.MTS -c:v libx265 -preset fast -crf 28 -tag:v hvc1 -c:a eac3 -b:a 224k output.mp4

exiftool -tagsFromFile "NAME.MTS" -time:all "output.mp4"
```

The first command converts the video and autio to a QuickTime-compatible format while the second one copies over the timestamp.  
To automate the process of converting all photos and preserving the folder structure, I wrote this Go script.
