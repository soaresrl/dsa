# Lista 5, 6, 7 - Estruturas de Dados - Heap - Sets - Disjoint Sets

## Árvore de arquivos

<pre>
├── lista_01
├── lista_03
├── lista_04
├── lista_05
    └── heaps.go
├── lista_06
    └── sets.go
├── lista_07
    └── partitions.go
├── README.md
└── pkg
    ├── bank/
    ├── circ_doubly_list/
    ├── circular_linked_list/
    ├── het_linked_list/
    ├── linked_list/
    ├── ord_doubly_list/
    ├── ordered_linked_list/
    ├── static_hash
    ├── extendible_hash
    ├── heap
        ├── heap.go <-- max-heap
        ├── min_heap.go
        └── heap_test.go
    ├── set
        ├── set.go
        └── set_test.go
    └── partition
        └── partition.go
</pre>

## Sobre a implementação

Foi implementada as Estrutura de _Max heap_ e Min heap_, _Sets_ e _Disjoint Sets_. Cada estrutura está em seu pacote (heap, set e partition, respectivamente).

## Programas referentes as Questões

Foram mantidos os arquivos referentes a Lista 01 e 04, de modo a manter o registro incremental das listas. A estrutura de dados de _heap_ com as operações solicitadas foram implementadas 
e estão incluídas na pasta `pkg/heap/heap.go` (_max heap_) e `pkg/heap/min_heap.go` (_min heap_). A estrutura de dados de _set_ com as operações solicitadas foram implementadas 
e estão incluídas na pasta `pkg/set/set.go`. A estrutura de dados de _disjoint set_ com as operações solicitadas foram implementadas 
e estão incluídas na pasta `pkg/partition/partition.go`.

O programa para execução contendo a função _main_ com as operações solicitadas para as estruturas de heap, set e partições encontra-se nas pastas `lista_05/heaps.go`, `lista_06/sets.go`, `lista_07/partitions.go`, respectivamente.

## Comandos

Para rodar a função _main_ relacionada a Lista 05 deve-se chamar o comando `go run lista_05/heaps.go`. Para rodar a função _main_ relacionada a Lista 06 deve-se chamar o comando `go run lista_06/sets.go`. Para rodar a função _main_ relacionada a Lista 07 deve-se chamar o comando `go run lista_07/partitions.go`.
