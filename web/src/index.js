const mainShowClass = 'fade-in-down';
const popupShowClass = 'popup-show';
const clipShowClass = 'clip-show';
const checkShowClass = 'check-show';
const mainHideClass = 'fade-out-down';
const durationClassPrefix = 'animate-';
const animateLinearClass = 'animate-linear';

const popupDurationMs = 700;

const animationTiming = {
    EASE: 'ease',
    LINEAR: 'linear',
};

/**
 * looks:
 * SPRINT NAME GEN FADES IN + OUT
 * CHOOSE A CATEGORY FADES IN
 * BUTTONS FADE IN
 *
 *
 * animations:
 * add all texts to queue and use a method which runs queue
 * overlap of animations should be possible, so that slot machine effect looks better
 *
 */

/**
 * todos:
 * - ---copy clipboard confirmation----
 * - buttons fade in with "select category"
 * - GitHub Link
 */

document.addEventListener('DOMContentLoaded', function () {
    if (!WebAssembly.instantiateStreaming) { // polyfill
        WebAssembly.instantiateStreaming = async (resp, importObject) => {
            const source = await (await resp).arrayBuffer();
            return await WebAssembly.instantiate(source, importObject);
        };
    }

    async function run() {
        const go = new Go();
        let result = await WebAssembly.instantiateStreaming(fetch("generator.wasm"), go.importObject);
        go.run(result.instance);
    }

    run().then(() => {

        /**
         * {
         *     text: string;
         *     animationDurationMs: number;
         *     animationTiming: "ease|linear";
         *     copyToClipboard: boolean;
         * }
         */
        let queue = [];
        let queueRunning = false;

        function addToQueue(textObjects) {
            if (!textObjects.length) {
                return;
            }
            queue = textObjects;
            runQueue();
        }

        async function runQueue(force) {
            if (queueRunning && force !== true) {
                return;
            }
            queueRunning = true;

            const next = queue[0];
            queue = queue.slice(1);

            await runAnimation(next);
            if (queue.length) {
                runQueue(true);
            } else {
                queueRunning = false;
                if (next.copyToClipboard) {
                    let currentEl = getCurrentTextEl();
                    currentEl.title = "Copy to Clipboard";
                    currentEl.addEventListener("click", copyToClipboard(currentEl), false);
                    currentEl.classList.add("pointer");
                }
            }
        }


        async function runAnimation(textObject) {
            let newTextEl = document.createElement('h1');
            newTextEl.classList.add('hidden');
            newTextEl.textContent = textObject.text;
            rootEl.appendChild(newTextEl);

            //hiding current text el
            let currentTextEl = getCurrentTextEl();
            currentTextEl.classList.add(`${durationClassPrefix}${textObject.animationDurationMs}`);
            if (textObject.animationTiming === animationTiming.LINEAR) {
                currentTextEl.classList.add(animateLinearClass)
            }
            currentTextEl.classList.add(mainHideClass);
            timeout(textObject.animationDurationMs).then(() => {
                currentTextEl.remove();
            });

            // showing new el
            newTextEl.classList.add(mainShowClass);
            newTextEl.classList.add(`${durationClassPrefix}${textObject.animationDurationMs}`);
            if (textObject.animationTiming === animationTiming.LINEAR) {
                newTextEl.classList.add(animateLinearClass)
            }
            newTextEl.classList.remove('hidden');
            await timeout(textObject.animationDurationMs);
            newTextEl.classList.remove(mainShowClass);
            newTextEl.classList.remove(`${durationClassPrefix}${textObject.animationDurationMs}`);
            newTextEl.classList.remove(animateLinearClass);

        }

        const rootEl = document.getElementById('textcontainer');

        function getCurrentTextEl() {
            return rootEl.children[1];
        }

        function timeout(ms) {
            return new Promise(resolve => {
                return setTimeout(resolve, ms);
            });
        }


        async function newSlotMachine(texts) {
            /* {
                *     text: string;
                *     animationDurationMs: number;
                *     animationTiming: "ease|linear";
                *     ?overlapMs: number;?
                * }
                */
            let textObjects = [];
            for (const text of texts) {
                textObjects.push({
                    text,
                    animationDurationMs: 200,
                    animationTiming: animationTiming.LINEAR,
                    copyToClipboard: true,
                });
            }
            addToQueue(textObjects);
        }

        function generateName(category) {
            let texts;
            switch (category) {
                case 'beer':
                    texts = [
                        "Heineken",
                        "Becks",
                        "Valley Lager",
                        "Mönchshofer",
                        "Heineken",
                        "Becks",
                        "ENDE",
                    ];
                    break;
                case 'movies':
                    texts = [
                        "Avengers",
                        "Titanic",
                        "Einer flog über das",
                        "Avengers",
                        "Titanic",
                        "Einer flog über das",
                    ];
                    break;
                case 'food':
                    texts = [
                        "Bread",
                        "Baguette",
                        "Roßwurst",
                        "Kaiserschmarrn",
                        "SUSHI",
                        "Braten",
                    ];
                    break;
                default:
                    texts = [];
            }

            return async () => {
                await newSlotMachine(texts);
            };
        }

        function generateNames(func) {
            return async () => {
                let names = [];
                for (let i = 0; i < 5; i++) {
                    names.push(func());
                }
                await newSlotMachine(names);
            }
        }

        document
            .getElementsByClassName("category animal")[0]
            .addEventListener("click", generateNames(NameGenerator.animal), false);
        document
            .getElementsByClassName("category docker")[0]
            .addEventListener("click", generateNames(NameGenerator.dockerName), false);
        document
            .getElementsByClassName("category marvel")[0]
            .addEventListener("click", generateNames(NameGenerator.marvelCharacter), false);
        document
            .getElementsByClassName("category dc")[0]
            .addEventListener("click", generateNames(NameGenerator.dcCharacter), false);
        document
            .getElementsByClassName("category game-of-thrones")[0]
            .addEventListener("click", generateNames(NameGenerator.got), false);


        let popupEl = document.getElementById('popup');

        function copyToClipboard(el) {
            return () => {
                navigator.clipboard.writeText(el.innerText).then(function () {
                    popupEl.classList.add(`${durationClassPrefix}${popupDurationMs}`);
                    popupEl.classList.add(popupShowClass);
                    popupEl.children[0].classList.add(`${durationClassPrefix}${popupDurationMs}`);
                    popupEl.children[0].classList.add(clipShowClass);
                    popupEl.children[1].classList.add(`${durationClassPrefix}${popupDurationMs}`);
                    popupEl.children[1].classList.add(checkShowClass);
                    timeout(popupDurationMs).then(() => {
                        popupEl.classList.remove(popupShowClass);
                        popupEl.children[0].classList.remove(clipShowClass);
                        popupEl.children[1].classList.remove(checkShowClass);
                    });
                }, function (err) {
                    console.error('Async: Could not copy text: ', err);
                });
            };
        }

        setTimeout( () => {
            if(!queue.length){
                addToQueue([
                    {
                        text: "PICK A CATEGORY",
                        animationDurationMs: 500,
                    },
                ]);
            }
        }, 1000);
    });
});