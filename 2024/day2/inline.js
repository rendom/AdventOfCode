// Part 1
document.getElementsByTagName("pre")[0].innerText.split("\n").filter(y => y).map(y => y.split(" ").map(z => Number(z)))
.filter((list) => {
    let dir = null;
    for(let i=1;i<list.length;i++) {
        const ldir = list[i-1] > list[i];
        if (dir === null) dir = ldir;
        if (ldir != dir) return false;
        if (Math.abs(list[i-1] - list[i]) < 1 || Math.abs(list[i-1] - list[i]) > 3) return false;
    }
    return true;
}).length
