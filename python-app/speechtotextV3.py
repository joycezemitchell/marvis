import speech_recognition as sr
import subprocess
import argparse

# Callback function for listen_in_background
def callback(recognizer, audio):
    try:
        text = recognizer.recognize_google(audio)
        print(f"Recognized: {text}")

        if 'hello' in text.lower():
            print("Trigger word 'hello' recognized.")
            # Save the recognized text to a file
            with open(filename, "w") as file:
                file.write(text)
            print(f"Text saved to {filename}")

    except sr.UnknownValueError:
        print("Could not understand the audio.")
    except sr.RequestError as e:
        print(f"Could not request results; check your network connection. Error: {e}")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Speech recognition script.')
    parser.add_argument('--filename', type=str, default='recognized_text.txt', help='Output filename for recognized text')
    args = parser.parse_args()

    filename = args.filename  # To make it accessible inside callback
    recognizer = sr.Recognizer()
    recognizer.dynamic_energy_threshold = True

    with sr.Microphone() as source:
        recognizer.adjust_for_ambient_noise(source, duration=1)
        print("Calibrated for ambient noise. Now, waiting for the trigger word 'hello'...")

        stop_listening = recognizer.listen_in_background(source, callback)

        # Keep the program running
        import time
        while True: time.sleep(0.1)