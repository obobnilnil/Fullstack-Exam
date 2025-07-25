from typing import List
from collections import Counter
import string

def find_most_frequent_word(messages: List[str], stop_words: List[str]) -> str:
    stop_set = set(word.lower() for word in stop_words)
    word_counter = Counter()

    for message in messages:
        for token in message.lower().split():
            word = token.strip(string.punctuation)
            if word and word not in stop_set:
                word_counter[word] += 1

    return max(word_counter, key=word_counter.get, default="")

if __name__ == "__main__":
    # messages = ["Hello world!", "Hello everyone", "The world is beautiful"]
    # stop_words = ["the", "is"]
    messages = ["The cat and the cat.", "A dog ran.", "The cat is fast."]
    stop_words = ["the", "a", "is", "and"]
    # stop_words = []
    # messages = ["covid19 is real", "Covid19!!! pandemic", "Covid19 covid19 covid19"]
    # stop_words = ["is"]

    result = find_most_frequent_word(messages, stop_words)
    print("Most frequent word:", result)
