let alreadyAuthd = false;

class Auth {
	static auth(service, token, cb) {
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
		.then(cb)
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
