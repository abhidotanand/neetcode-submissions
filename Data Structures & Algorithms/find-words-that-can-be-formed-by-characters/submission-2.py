from collections import Counter
class Solution:
    def countCharacters(self, words: List[str], chars: str) -> int:
        charFreq = Counter(chars)
        total = 0

        for word in words:
            wordFreq = Counter(word)

            # if wordFreq - charFreq is empty, the word is good!
            # ex: word = ate, chars = bate. word-chars = ate(3) - bate(4) = -1
                # => Counter floors to 0 aka empty!
            if not (wordFreq - charFreq):
                total += len(word)

        return total

        # time complexity: O(N * K + M), where N = words, k = chars of each word, m = chars
        # space complexity: O(K + M)