import sys
from gtts import gTTS

# Check if the correct number of arguments is passed
if len(sys.argv) != 3:
    print("Usage: python text_to_speech.py input_text_file.txt output_audio_file.mp3")
    sys.exit(1)

# Assigning command line arguments to variables
input_text_file = sys.argv[1]
output_audio_file = sys.argv[2]

# Reading the text from the file
with open(input_text_file, 'r', encoding='utf-8') as file:
    text_to_say = file.read()

# Using gTTS to convert the text to speech
tts = gTTS(text=text_to_say, lang='en')

# Saving the converted speech to an MP3 file
tts.save(output_audio_file)




