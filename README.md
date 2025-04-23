

## Real-World Example:

- Suppose a customer enters:

{
  "name": "Johnathan Smith",
  "address": "123 Main St, London"
}

- If the Operator has:

    {
    "name": "Jonathan Smith",
    "address": "123 Main Street, London"
    }

- The Jaro-Winkler score might return something like:

{
  "matchScore": 94
}



==================

## What is Jaro-Winkler Algorithem

- Jaro-Winkler is a string similarity algorithm that measures how closely two strings resemble each other. It gives a score between 0 and 1 (or 0 to 100 if scaled) — with 1 or 100 meaning an exact match.

It's commonly used in:

        - Identity verification (like KYC)

        - Data deduplication

        - Spell-checking

        - Record linkage


## How It Works
- Jaro-Winkler is built in two parts:

1. Jaro Distance
This is the base similarity score between two strings. It considers:

- Matching characters

- Number of transpositions (order differences)

2. Winkler Boost
- If the two strings share a common prefix, the score is slightly boosted to reflect the likelihood of it being a good match. This is useful in names and identity matching  (e.g., “Jon” vs “John”).

 ## Example Comparison
Let’s compare:

String A: MARTHA
String B: MARHTA

Matching characters: M, A, R, T, H, A (6 characters)

Transpositions: H and T are swapped → 1 transposition

🔹 Jaro score ≈ 0.944
Then since they share a common prefix MAR, Winkler boost is applied, giving

🔸 Jaro-Winkler score ≈ 0.961


 ## Formula

    J = 1/3 (m/s1 + m/s2 + m-t/m)

   ## Where:

   m = number of matching characters
   t = number of transpositions ÷ 2
   s1,s2 = two strings being compared

   JW Distance : JW = J + (l.p.(1 - J))
        l = length of common prefix
        p = scalling factor



 