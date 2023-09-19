local array = {}

for i = 1, 3 do
    array[i] = {}
    for j = 1, 3 do
        array[i][j] = i * j
    end
end

print(array)

debug()