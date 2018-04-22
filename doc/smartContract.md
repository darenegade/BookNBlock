# Smart Contract
## Generelle Idee
* Für jede Wohnung gibt es eine [State Machine](https://solidity.readthedocs.io/en/v0.4.21/common-patterns.html#state-machine "")
* Aktionen des Mieters/Vermieters ändern diesen Zustand
* Zeitgesteuerte Aktionen sind auch möglich

## Zustände
1. buchbar
2. reserviert
3. gebucht

## Übergänge
* -> 1: Wohnung anlegen
* 1 -> 2: Mieter frägt Wohnung an
* 1 -> 3: Vermieter nutzt Wohnung selbst
* 2 -> 3: Vermieter stimmt Anfrage zu
* 2 -> 1: Vermieter lehnt Anfrage ab
* 3 -> 1: Stornierung durch Mieter
* 3 -> 1: Eigennutzung anmelden