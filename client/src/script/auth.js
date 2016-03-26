let alreadyAuthd = false;
let rootComponent = null;

class Auth {
	static setRootComponent(rc) {
		rootComponent = rc;
	}

	static auth(service, token) {
		fetch('/auth', {
			method: 'POST',
			credentials: 'same-origin',
			body: JSON.stringify({
				service: service,
				token: token,
			})
		})
		.then(function() {
			alreadyAuthd = true;
		})
		.then(function() {
			if (rootComponent) {
				rootComponent.forceUpdate();
			}
		})
		.catch(function (error) {
			console.log(error);
		});
	}

	static isAuthd() {
		if (alreadyAuthd) {
			return true;
		}
		let jwt = getCookieValue("auth");
		if (jwt) {
			alreadyAuthd = true;
		}
		return alreadyAuthd;
	}
}

function getCookieValue(name) {
	let parsed = document.cookie.match('(^|;)\\s*' + name + '\\s*=\\s*([^;]+)');
	return parsed ? parsed.pop() : '';
}

module.exports = Auth
