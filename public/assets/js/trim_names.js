function trimNames(className, maxLength, split = false) {
    const elements = document.getElementsByClassName(className);
    for (const element of elements) {
        let text = element.innerText;
        if (text.length > maxLength) {
            if (split) {
                text = text.split(" ")[1];
            }
            let original = text;
            element.title = original;
            text = text.substr(0, maxLength - 3) + "…";
            element.innerHTML = element.innerHTML.replace(original, text);
        }
    }
}
trimNames("player_name", 15, true);
trimNames("clan_name", 20);
