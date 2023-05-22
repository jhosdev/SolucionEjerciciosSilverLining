//usamos reduce() para sumar los elementos de un array sin usar un bucle for, complejidad O(n)

function sumN(arr) {
    return arr.reduce((acumulador, numero) => acumulador + numero, 0);
}

const cant = 10;
const num = Array.from({ length: cant }, () => Math.floor(Math.random() * 20));

console.log(num);

const res = sumN(num);

console.log(res);