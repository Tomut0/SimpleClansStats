import {parse} from 'yaml';
import {toHex} from "@/colors.js";

const bannerImagesMap = new Map();

async function createBanner(yaml, urls, ctx, width = 20, height = 40) {
    const images = getImages(urls);
    const sectors = sectorize(yaml);

    for (const [key, color] of sectors) {
        const pattern = getPattern(key);
        if (pattern === "unknown") {
            console.warn(`Can't find a pattern for key ${key}.`);
        }

        const image = await images.get(getPattern(key));
        ctx.drawImage(drawByColor(image.target, width, height, color), 0, 0);
    }
}

/**
 * @param code minecraft's code
 * @returns {string} resource name from code
 * @see https://minecraft.wiki/w/Banner/Patterns
 */
const getPattern = (code) => {
    switch (code) {
        case "b":
            return "base";
        case "bs":
            return "stripe_bottom";
        case "ts":
            return "stripe_top";
        case "ls":
            return "stripe_left";
        case "rs":
            return "stripe_right";
        case "ms":
            return "stripe_center";
        case "drs":
            return "stripe_downright";
        case "dls":
            return "stripe_downleft";
        case "ss":
            return "small_stripes";
        case "cr":
            return "cross";
        case "sc":
            return "straight_cross";
        case "ld":
            return "diagonal_left";
        case "rd":
            return "diagonal_up_right";
        case "rud":
            return "diagonal_right";
        case "lud":
            return "diagonal_up_left";
        case "vh":
            return "half_vertical";
        case "vhr":
            return "half_vertical_right";
        case "hh":
            return "half_horizontal";
        case "hhb":
            return "half_horizontal_bottom";
        case "bl":
            return "square_bottom_left";
        case "br":
            return "square_bottom_right";
        case "tl":
            return "square_top_left";
        case "tr":
            return "square_top_right";
        case "bt":
            return "triangle_bottom";
        case "tt":
            return "triangle_top";
        case "bts":
            return "triangles_bottom";
        case "tts":
            return "triangles_top";
        case "mc":
            return "circle";
        case "mr":
            return "rhombus";
        case "bo":
            return "border";
        case "cbo":
            return "curly_border";
        case "bri":
            return "bricks";
        case "gra":
            return "gradient";
        case "gru":
            return "gradient_up";
        case "cre":
            return "creeper";
        case "sku":
            return "skull";
        case "flo":
            return "flower";
        case "moj":
            return "mojang";
        case "glb":
            return "globe";
        case "pig":
            return "piglin";
        case "flw":
            return "flow";
        case "gus":
            return "guster";
        default:
            return "unknown";
    }
};

const drawByColor = (image, width, height, rgbColor) => {
    const canvas = document.createElement('canvas');
    const ctx = canvas.getContext('2d');

    ctx.drawImage(image, 1, 1, 20, 40, 0, 0, width, height);

    const imageData = ctx.getImageData(0, 0, width, height);

    colorImage(imageData, rgbColor);

    ctx.putImageData(imageData, 0, 0);

    return canvas;
}

const colorImage = (imageData, rgbColor) => {
    const data = imageData.data;
    for (let i = 0; i < data.length; i += 4) {
        const brightness = Math.floor(0.2 * data[i] + 0.7 * data[i + 1] + 0.1 * data[i + 2]);

        data[i] = rgbColor[0];
        data[i + 1] = rgbColor[1];
        data[i + 2] = rgbColor[2];
        data[i + 3] = brightness;
    }
}

function getImages(urls) {
    if (bannerImagesMap.size > 0) {
        return bannerImagesMap;
    }

    for (const url of urls) {
        const img = new Image();
        img.src = url;

        const key = url.split("/").pop().split('.')[0];
        const value = new Promise((resolve) => img.addEventListener("load", resolve));

        bannerImagesMap.set(key, value);
    }

    return bannerImagesMap;
}

function getColorFromType(type) {
    const base = type.split("_", 3);

    // example: BLUE_BANNER
    if (base.length === 2) {
        return toHex(base[0]);
    }

    // example: LIGHT_BLUE_BANNER
    if (base.length === 3) {
        return toHex(`${base[0]}_${base[1]}`);
    }

    throw new TypeError("Invalid banner type");
}

function isNotBanner(obj) {
    if (typeof obj !== 'object') {
        obj = parse(obj);
    }

    return !obj ||
        !Object.hasOwn(obj, 'cs') ||
        !Object.hasOwn(obj.cs, 'type') ||
        !obj.cs.type.includes("BANNER");
}

function sectorize(yaml) {
    const sectors = new Map();
    const banner = parse(yaml);

    if (isNotBanner(banner)) {
        throw new TypeError("Invalid banner object");
    }

    sectors.set("b", hexToRGB(getColorFromType(banner.cs.type)));

    if (Object.hasOwn(banner.cs, "meta") && Object.hasOwn(banner.cs.meta, "patterns")) {
        for (const pattern of banner.cs.meta.patterns) {
            sectors.set(pattern['pattern'], hexToRGB(toHex(pattern['color'])));
        }
    }

    return sectors;
}

function hexToRGB(hex, alpha = 1) {
    const r = parseInt(hex.slice(1, 3), 16),
        g = parseInt(hex.slice(3, 5), 16),
        b = parseInt(hex.slice(5, 7), 16);

    return [r, g, b, alpha];
}

export {createBanner, isNotBanner};

