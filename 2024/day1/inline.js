// Part 1
const list=[[],[]]; document.getElementsByTagName("pre")[0].innerText.split("\n").filter(x => x).forEach((x) => { let a = x.split("   ").map((y) => Number(y)); list[0].push(a[0]); list[1].push(a[1]) }); list[0].sort(); list[1].sort(); list[0].reduce((a, curr, idx) => Math.abs(curr - list[1][idx]) + a, 0);

// Part 2
const list=[[],[]]; document.getElementsByTagName("pre")[0].innerText.split("\n").filter(x => x).forEach((x) => { let a = x.split("   ").map((y) => Number(y)); list[0].push(a[0]); list[1].push(a[1]) }); list[0].reduce((a, curr) => (curr * list[1].filter(y => y == curr).length) + a)
