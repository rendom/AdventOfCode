// Part 1
document.querySelector("pre").innerText.split("\n").filter(x => x).reduce((a, b) => { let m = b.match(/\d/g); return Number(m[0]+m[m.length-1])+a; }, 0);

// Part 2
const z = ['zero', 'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine']; document.querySelector("pre") .innerText.split("\n") .filter(x => x) .reduce((a, b) => { const m = Array.from(b.matchAll(new RegExp(`(?=(${z.join('|')}|[0-9]))`,'g'))).map((m1) => { return m1[1].replaceAll(new RegExp(z.join('|'), 'g'), (a) => { const i = z.indexOf(a); return i == -1 ? a : i }) }); return Number(m[0]+m[m.length-1])+a; }, 0);
