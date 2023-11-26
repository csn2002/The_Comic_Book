const prevBtn = document.querySelector('#prev-btn');
const nextBtn = document.querySelector('#next-btn');
const book = document.querySelector('#book');
const papers = document.querySelectorAll('.paper');
prevBtn.addEventListener("click", goPrevious);
nextBtn.addEventListener("click", goNext);

// Business Logic
let currentState =1;
let numOfPapers = papers.length+1;
let maxState = numOfPapers + 1;


function openBook() {
	book.style.transform = "translateX(50%)";
	prevBtn.style.transform = "translateX(-180px)";
	nextBtn.style.transform = "translateX(180px)";
}

function closeBook(isFirstPage) {
	if(isFirstPage) {
		book.style.transform = "translateX(0%)";
	} else {
		book.style.transform = "translateX(100%)";
	}
	prevBtn.style.transform = "translateX(0px)";
	nextBtn.style.transform = "translateX(0px)";
}

function goNext() {
	console.log(currentState)
	if(currentState < maxState-1) {
		if(currentState == 1){
			openBook();
			papers[currentState - 1].classList.add("flipped");
			papers[currentState - 1].style.zIndex = currentState;
		}
		else if(currentState==maxState-2){
			closeBook(false);
			papers[currentState - 1].classList.add("flipped");
			papers[currentState - 1].style.zIndex = currentState;
		}
		else{
			papers[currentState - 1].classList.add("flipped");
			papers[currentState - 1].style.zIndex = currentState;
		}
		currentState++;
	}
	else{
		throw new Error("unkown state");
	}
}

function goPrevious() {
	console.log(currentState)
	if(currentState > 1) {
		if(currentState==2){
			closeBook(true);
			papers[currentState - 2].classList.remove("flipped");
			papers[currentState - 2].style.zIndex = maxState - currentState+1;
		}
		else if(currentState==maxState-1){
			openBook()
			papers[currentState - 2].classList.remove("flipped");
			papers[currentState - 2].style.zIndex = maxState - currentState+1;
		}
		else{
			papers[currentState - 2].classList.remove("flipped");
			papers[currentState - 2].style.zIndex = maxState - currentState+1;
			// break;
		}
		currentState--;
	}
	else{
		throw new Error("unkown state");
	}
}
var navLinks = document.querySelectorAll("nav a");
for (var i = 0; i < navLinks.length; i++) {
	var link = navLinks[i]
	if (link.getAttribute('href') == window.location.pathname) {
		link.classList.add("live");
		break;
	}
}
function createInputFields() {
	// Clear the previous input fields
	document.getElementById('inputFieldsContainer').innerHTML = '';

	// Get the number of inputs from the user
	const numInputs = document.getElementById('numInputs').value;

	// Create new input fields based on the user input
	for (let i = 0; i < numInputs; i++) {
		const inputField = document.createElement('input');
		inputField.type = 'text';
		inputField.name = `input_${i + 1}`;
		inputField.placeholder = `Enter string ${i + 1}`;

		// Append the input field to the container
		document.getElementById('inputFieldsContainer').appendChild(inputField);
	}
}

