import speech_recognition as sr
import subprocess
import argparse

# Function to recognize speech and save to a text file
def recognize_and_save(filename):
    recognizer = sr.Recognizer()

    while True:  # Infinite loop
        with sr.Microphone() as source:
            print("Waiting for the trigger word 'hello'...")
            audio_data = recognizer.listen(source)
            try:
                # Recognize the speech using Google's speech recognition
                text = recognizer.recognize_google(audio_data)
                print("You said:", text)

                # Save the recognized text to a file
                with open(filename, "w") as file:
                    file.write(text)
                print(f"Text saved to {filename}")

                # Play sound using ffplay
                subprocess.run(["ffplay", "-v", "0", "-nodisp", "-autoexit", "ding.mp3"], check=True)

                break  # Exit the infinite loop

            except sr.UnknownValueError:
                print("Could not understand the audio.")
            except sr.RequestError:
                print("Could not request results; check your network connection.")
            except Exception as e:
                print(f"An error occurred: {e}")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Speech recognition script.')
    parser.add_argument('--filename', type=str, default='recognized_text.txt', help='Output filename for recognized text')
    args = parser.parse_args()

    recognize_and_save(args.filename)