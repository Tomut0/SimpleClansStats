import './bootstrap';
import '../css/app.css';
import "/node_modules/flag-icons/css/flag-icons.min.css";
import 'chartjs-adapter-moment';

import {createApp, h} from 'vue';
import {createInertiaApp} from '@inertiajs/vue3';
import {resolvePageComponent} from 'laravel-vite-plugin/inertia-helpers';
import {ZiggyVue} from '../../vendor/tightenco/ziggy/dist/vue.m';
import { MotionPlugin } from '@vueuse/motion'

const appName = import.meta.env.VITE_APP_NAME || 'Laravel';

createInertiaApp({
    title: (title) => `${title} - ${appName}`,
    resolve: (name) => resolvePageComponent(`./Pages/${name}.vue`, import.meta.glob('./Pages/**/*.vue')),
    setup({el, App, props, plugin}) {
        return createApp({render: () => h(App, props)})
            .use(plugin)
            .provide('appName', appName)
            .use(ZiggyVue, Ziggy)
            .use(MotionPlugin)
            .mount(el);
    },
    progress: {
        color: '#4B5563',
    },
});
