var x = document.querySelector("pre").innerText.split("\n\n").map((x) => x.split("\n").reduce((a,b) => { return Number(a) + Number(b)}, 0)).sort((a,b) => b-a); 
console.log(`top: ${x[0]} top3total: ${x.slice(0,3).reduce((a,b) => { return Number(a) + Number(b)}, 0)}`);
