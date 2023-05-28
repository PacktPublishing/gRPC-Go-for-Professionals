# Chapter 2 - Protobuf Primer

These text files in this folder represent the values that we want to serialize. With these we can create an instance of `Account` (defined in `account.proto`) or Wrappers (defined in `wrappers.proto`), and encode data to binary.

An example of encoding is the following:

## Linux/Mac
```shell
cat account.txt | protoc --encode=Account account.proto | hexdump -C
```

## Windows (Powershell)

```shell
(Get-Content account.txt | protoc --encode=Account account.proto) -join "`n" | Format-Hex
```