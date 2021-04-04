function filter(rows, text) {
    outter:
        for (const row of rows) {
            for (const column of row.children) {
                if (column.innerText.toLowerCase().includes(text.toLowerCase())) {
                    row.hidden = false;
                    continue outter;
                }
            }
            row.hidden = true;
        }
}

const filterInput = document.getElementById("filter");
const listener = function () {
    let text = filterInput.value;
    const rows = document.getElementsByClassName("data_row");
    filter(rows, text);
}
filterInput.addEventListener("input", listener);
window.addEventListener("load", listener);
