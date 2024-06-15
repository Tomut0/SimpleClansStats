/**
 * Get the text color based on the background color.
 *
 * @author SudoPlz
 * @link https://stackoverflow.com/a/41491220/24461279
 *
 * @param bgColor
 * @param lightColor
 * @param darkColor
 * @returns {string}
 */
function getColorByBgColor(bgColor, lightColor = '#FFF', darkColor = '#000') {
    const color = (bgColor.charAt(0) === '#') ? bgColor.substring(1, 7) : bgColor;
    const r = parseInt(color.substring(0, 2), 16); // hexToR
    const g = parseInt(color.substring(2, 4), 16); // hexToG
    const b = parseInt(color.substring(4, 6), 16); // hexToB
    return (((r * 0.299) + (g * 0.587) + (b * 0.114)) > 186) ?
        darkColor : lightColor;
}

/**
 * Extract the first color from a colored clan tag
 *
 * @example
 * <pre>
 *      extractFirstColor("[&aClan&r] Player") // returns "#55FF55"
 *      extractFirstColor("[&#00FF00Clan&r] Player") // returns "#00FF00"
 * </pre>
 * @param coloredTag
 * @returns {string}
 */
function extractFirstColor(coloredTag) {
    // Regular expression to find Minecraft color codes
    const colorCodeRegex = /(&#[0-9A-Fa-f]{6}|&[0-9A-Fa-f])/;
    const match = coloredTag.match(colorCodeRegex);

    if (match) {
        const colorCode = match[0];
        if (colorCode.startsWith("&#")) {
            // If the color is in hexadecimal format
            return colorCode.slice(1); // Remove the initial '&'
        } else {
            // If the color is in Minecraft code format
            switch (colorCode[1].toLowerCase()) {
                case '0':
                    return '#000000'; // black
                case '1':
                    return '#0000AA'; // dark blue
                case '2':
                    return '#00AA00'; // dark green
                case '3':
                    return '#00AAAA'; // dark aqua
                case '4':
                    return '#AA0000'; // dark red
                case '5':
                    return '#AA00AA'; // dark purple
                case '6':
                    return '#FFAA00'; // gold
                case '7':
                    return '#AAAAAA'; // gray
                case '8':
                    return '#555555'; // dark gray
                case '9':
                    return '#5555FF'; // blue
                case 'a':
                    return '#55FF55'; // green
                case 'b':
                    return '#55FFFF'; // aqua
                case 'c':
                    return '#FF5555'; // red
                case 'd':
                    return '#FF55FF'; // pink
                case 'e':
                    return '#FFFF55'; // yellow
                case 'f':
                    return '#FFFFFF'; // white
                default:
                    return '#FFFFFF'; // unrecognized color code
            }
        }
    }
    return '#FFFFFF'; // if no matches
}

/**
 * Adds colors to a string based on Minecraft-style color codes.
 *
 * @author lpostiglione
 * @link https://github.com/lpostiglione/SimpleClansStats2/blob/master/includes/functions.inc.php
 * @param {string} string - The input string without colors.
 * @return {string} - The string with colors.
 */
function addColors(string) {
    const colors = {
        "g": "<span class=\"color-black\">",
        "1": "<span class=\"color-dark_blue\">",
        "2": "<span class=\"color-dark_green\">",
        "3": "<span class=\"color-dark_aqua\">",
        "4": "<span class=\"color-dark_red\">",
        "5": "<span class=\"color-dark_purple\">",
        "6": "<span class=\"color-gold\">",
        "7": "<span class=\"color-gray\">",
        "8": "<span class=\"color-dark_gray\">",
        "9": "<span class=\"color-blue\">",
        "a": "<span class=\"color-green\">",
        "b": "<span class=\"color-aqua\">",
        "c": "<span class=\"color-red\">",
        "d": "<span class=\"color-light_purple\">",
        "e": "<span class=\"color-yellow\">",
        "f": "<span class=\"color-white\">",
        "l": "<span class=\"bold\">",
        "m": "<span class=\"strikethrough\">",
        "n": "<span class=\"underline\">",
        "o": "<span class=\"italic\">",
        "k": "<span class=\"obfuscated\">"
    };

    string = string.replace(/§0/g, "§g");
    const motdarr = string.split("§");
    let colored = '<span class="colored">';
    let openSpans = 0;

    for (const row of motdarr) {
        if (row === "") continue;

        const colorCode = row[0].toLowerCase();
        const text = row.slice(1);

        if (colorCode === "r") {
            colored += "</span>".repeat(openSpans);
            openSpans = 0;
        } else if (colors[colorCode]) {
            colored += "</span>".repeat(openSpans);
            openSpans = 0;
            colored += colors[colorCode];
            openSpans++;
        } else {
            colored += row;
        }

        colored += text;
    }

    colored += "</span>".repeat(openSpans);
    return colored + "</span>";
}

export {getColorByBgColor, extractFirstColor, addColors};
