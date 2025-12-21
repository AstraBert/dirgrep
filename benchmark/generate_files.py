import random as rand

RANDOM_CHUNKS = [
    "A lonely kite danced above silent rooftops.",
    "Coffee cooled while the sunrise argued with clouds.",
    "Unexpected laughter spilled from the empty hallway.",
    "Her notebook whispered plans to tomorrow.",
    "Rain tapped messages onto impatient windows.",
    "The old clock practiced forgetting time.",
    "Bright socks negotiated peace with formal shoes.",
    "A cat supervised the moonrise seriously.",
    "Wind borrowed names from passing trains.",
    "Paper airplanes dreamt of permanent addresses.",
    "The library exhaled dust and patience.",
    "Morning emails queued politely for attention.",
    "Stars rehearsed excuses behind thin clouds.",
    "A button escaped, declaring independence midlaundry.",
    "The map blushed, embarrassed by shortcuts.",
    "Silence negotiated terms after the thunder.",
    "Two pigeons debated philosophy over crumbs.",
    "The elevator hummed secrets between floors.",
    "Autumn rehearsed colors in quiet parks.",
    "A password forgot itself at dawn."
]

def generate_file(i: int) -> None:
    chunks = []
    for _ in range(3):
        chunks.append(RANDOM_CHUNKS[rand.randint(0, len(RANDOM_CHUNKS)-1)])
    with open(f"files/file{i}.txt", "w") as f:
        f.writelines(chunks)
    return None

if __name__ == "__main__":
    import os
    import sys

    print("This program is about to create 1.000.000 files (taking approx. 4GB of space): are you sure you want to proceed? [y/n]")
    answer = input(">>> ")

    if answer.strip().lower() in ("yes", "y", "yse"):
        os.makedirs("files/", exist_ok=True)
        for i in range(1000000):
            if (i+1)%100000 == 0:
                print(f"Processed {i+1} files")
            generate_file(i+1)
        sys.exit(0)
    else:
        sys.exit(1)
    