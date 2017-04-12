import { sayHello } from './greet';

// Write a greeting in the element
const element = document.getElementById('greeting');
element.innerHTML = sayHello("Django seats");
