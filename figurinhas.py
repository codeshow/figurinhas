# %%
import random

N = 500             # Tamanho do álbum
n = 4               # Tamanho do pacote
total_albuns = 1000 # Quantidade de Simulações

def simula_album(N, n):
    album = set()
    pacotes = 0
    while len(album) < N:
        pacote = [random.randint(1, N) for i in range(n)]
        album = album.union(pacote)
        pacotes += 1
    return pacotes

resultados = [simula_album(N, n) for i in range(total_albuns)]
print("Média de pacotes a serem comprados:", sum(resultados) / len(resultados))