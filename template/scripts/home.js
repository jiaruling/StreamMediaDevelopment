// note: 初始化
$(document).ready(function() {

	DEFAULT_COOKIE_EXPIRE_TIME = 30; // cookie 过期时间

	uname = '';
	session = '';
	uid = 0;
	currentVideo = null;
	listedVideos = null;

	session = getCookie('session');
	uname = getCookie('username');

	initPage(function() {
		if (listedVideos !== null) {
			currentVideo = listedVideos[0];
			selectVideo(listedVideos[0]['id']);
		}

		$(".video-item").click(function() {
			var self = this.id
  			listedVideos.forEach(function(item, index) {
  				if (item['id'] === self) {
  					currentVideo = item;
  					return
  				}
  			});

  			selectVideo(self);
		});

		$(".del-video-button").click(function() {
			var id = this.id.substring(4);
  			deleteVideo(id, function(res, err) {
  				if (err !== null) {
  					//window.alert("encounter an error when try to delete video: " + id);
  					popupErrorMsg("encounter an error when try to delete video: " + id);
  					return;
  				}

  				popupNotificationMsg("Successfully deleted video: " + id)
  				location.reload();
  			});
		});

		$("#submit-comment").on('click', function() {
			var content = $("#comments-input").val();
  			postComment(currentVideo['id'], content, function(res, err) {
  				if (err !== null) {
  					popupErrorMsg("encounter and error when try to post a comment: " + content);
  					return;
  				}

  				if (res === "ok") {
  					popupNotificationMsg("New comment posted")
    				$("#comments-input").val("");

    				refreshComments(currentVideo['id']);
  				}
  			});
		});
	});

	// home page event registry
	$("#regbtn").on('click', function(e) {
		$("#regbtn").text('Loading...')
    		e.preventDefault()
    		registerUser(function(res, err) {
    			if (err != null) {
    				$('#regbtn').text("Register")
    				popupErrorMsg('encounter an error, pls check your username or pwd');
    				return;
    			}

    			var obj = JSON.parse(res);
    			setCookie("session", obj["session_id"], DEFAULT_COOKIE_EXPIRE_TIME);
    			setCookie("username", uname, DEFAULT_COOKIE_EXPIRE_TIME);
    			$("#regsubmit").submit();
    		});
	});

	$("#siginbtn").on('click', function(e) {

		$("#siginbtn").text('Loading...')
    	e.preventDefault();
    	signinUser(function(res, err) {
    		if (err != null) {
    			$('#siginbtn').text("Sign In");
    		//window.alert('encounter an error, pls check your username or pwd')
    			popupErrorMsg('encounter an error, pls check your username or pwd');
    			return;
    		}

    		var obj = JSON.parse(res);
    		setCookie("session", obj["session_id"], DEFAULT_COOKIE_EXPIRE_TIME);
    		setCookie("username", uname, DEFAULT_COOKIE_EXPIRE_TIME);
    		$("#siginsubmit").submit();
    	});
	});

	$("#signinhref").on('click', function() {
		$("#regsubmit").hide();
		$("#siginsubmit").show();
	});

	$("#registerhref").on('click', function() {
		$("#regsubmit").show();
		$("#siginsubmit").hide();
	});

	// userhome event register
	$("#upload").on('click', function() {
  		$("#uploadvideomodal").show();

  	});

	$("#uploadform").on('submit', function(e) {
		e.preventDefault()
	  	var vname = $('#vname').val();

	  	createVideo(vname, function(res, err) {
	  		if (err != null ) {
	  			//window.alert('encounter an error when try to create video');
	  			popupErrorMsg('encounter an error when try to create video');
	  			return;
	  		}

	  		var obj = JSON.parse(res);
	  		var formData = new FormData();
			formData.append('file', $('#inputFile')[0].files[0]);

			$.ajax({
				url : 'http://' + window.location.hostname + ':8000/upload/' + obj['id'],
				//url:'http://127.0.0.1:8000/upload/dbibi',
				type : 'POST',
				data : formData,
				//headers: {'Access-Control-Allow-Origin': 'http://127.0.0.1:9000'},
				crossDomain: true,
				processData: false,  // tell jQuery not to process the data
				contentType: false,  // tell jQuery not to set contentType
				success : function(data) {
				   console.log(data);
				   $('#uploadvideomodal').hide();
				   location.reload();
				   //window.alert("hoa");
				},
				complete: function(xhr, textStatus) {
					if (xhr.status === 204) {
						window.alert("finish")
						return;
					}
					if (xhr.status === 400) {
						$("#uploadvideomodal").hide();
						popupErrorMsg('file is too big');
						return;
					}
				}
			});
	  	});
	});

	$(".close").on('click', function() {
		$("#uploadvideomodal").hide();
	});

	$("#logout").on('click', function() {
		setCookie("session", "", -1)
		setCookie("username", "", -1)
	});


    $(".video-item").click(function () {
  	    var url = 'http://' + window.location.hostname + ':9000/videos/'+ this.id
  	    var video = $("#curr-video");
  	    video[0].attr('src', url);
  	    video.load();
    });
});

// note: 初始化页面
function initPage(callback) {
	getUserId(function(res, err) {
		if (err != null) {
			window.alert("Encountered error when loading user id");
			return;
		}

		var obj = JSON.parse(res);
		uid = obj['id'];
		//window.alert(obj['id']);
		listAllVideos(function(res, err) {
			if (err != null) {
				//window.alert('encounter an error, pls check your username or pwd');
				popupErrorMsg('encounter an error, pls check your username or pwd');
				return;
			}
			var obj = JSON.parse(res);
			listedVideos = obj['videos'];
			obj['videos'].forEach(function(item, index) {
				var ele = htmlVideoListElement(item['id'], item['name'], item['display_ctime']);
				$("#items").append(ele);
			});
			callback();
		});
	});
}

// note: session
function setCookie(cname, cvalue, exmin) {
    var d = new Date();
    d.setTime(d.getTime() + (exmin * 60 * 1000));
    var expires = "expires="+d.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

function getCookie(cname) {
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for(var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

// note: DOM operations
function selectVideo(vid) {
	var url = 'http://' + window.location.hostname + ':8000/videos/'+ vid
  	var video = $("#curr-video");
  	$("#curr-video:first-child").attr('src', url);
  	$("#curr-video-name").text(currentVideo['name']);
  	$("#curr-video-ctime").text('Uploaded at: ' + currentVideo['display_ctime']);
  	//currentVideoId = vid;
  	refreshComments(vid);
}

function refreshComments(vid) {
	listAllComments(vid, function (res, err) {
		if (err !== null) {
			//window.alert("encounter an error when loading comments");
			popupErrorMsg('encounter an error when loading comments');
			return
		}

		var obj = JSON.parse(res);
		$("#comments-history").empty();
		if (obj['comments'] === null) {
			$("#comments-total").text('0 Comments');
		} else {
			$("#comments-total").text(obj['comments'].length + ' Comments');
		}
		obj['comments'].forEach(function(item, index) {
			var ele = htmlCommentListElement(item['id'], item['user'], item['content']);
			$("#comments-history").append(ele);
		});

	});
}

function popupNotificationMsg(msg) {
	var x = document.getElementById("snackbar");
	$("#snackbar").text(msg);
    x.className = "show";
    setTimeout(function(){ x.className = x.className.replace("show", ""); }, 2000);
}

function popupErrorMsg(msg) {
	var x = document.getElementById("errorbar");
	$("#errorbar").text(msg);
    x.className = "show";
    setTimeout(function(){ x.className = x.className.replace("show", ""); }, 2000);
}

function htmlCommentListElement(cid, author, content) {
	var ele = $('<div/>', {
		id: cid
	});

	ele.append(
		$('<div/>', {
		  class: 'comment-author',
		  text: author + ' says:'
		})
	);
	ele.append(
		$('<div/>', {
		  class: 'comment',
		  text: content
		})
	);

	ele.append('<hr style="height: 1px; border:none; color:#EDE3E1;background-color:#EDE3E1">');

    return ele;
}

function htmlVideoListElement(vid, name, ctime) {
	var ele = $('<a/>', {
		href: '#'
	});
	ele.append(
		$('<video/>', {
			width:'320',
			height:'240',
			poster:'/statics/img/preloader.jpg',
			controls: true
			//href: '#'
		})
	);
	ele.append(
		$('<div/>', {
			text: name
		})
	);
	ele.append(
		$('<div/>', {
			text: ctime
		})
	);

	var res = $('<div/>', {
		id: vid,
		class: 'video-item'
	}).append(ele);

	res.append(
		$('<button/>', {
			id: 'del-' + vid,
			type: 'button',
			class: 'del-video-button',
			text: 'Delete'
		})
	);

	res.append(
		$('<hr>', {
			size: '2'
		}).css('border-color', 'grey')
	);

	return res;
}

// note: Async ajax methods

// note: User operations
// DONE: 完成
function registerUser(callback) {
	var username = $("#username").val();
	var pwd = $("#pwd").val();

	if (username == '' || pwd == '') {
		callback(null, err);
	}

	var reqBody = {
		'username': username,
		'pwd': pwd
	}

	var dat = {
		'url': 'http://127.0.0.1:8001/user',
		'method': 'POST',
		'req_body': JSON.stringify(reqBody)
	};

	$.ajax({
		url  : '/api',
		type : 'post',
		data : JSON.stringify(dat),
		statusCode: {
			500: function() {
				callback(null, "internal error");
			}
		},
		complete: function(xhr, textStatus) {
			if (xhr.status >= 400) {
				callback(null, "Error of Signin");
				return;
			}
		}
	}).done(function(data, statusText, xhr){
		if (xhr.status >= 400) {
			callback(null, "Error of register");
			return;
		}

		uname = username;
		callback(data, null);
	});
}

// DONE: 完成
function signinUser(callback) {
	var username = $("#susername").val();
	var pwd = $("#spwd").val();
	if (username == '' || pwd == '') {
		callback(null, err);
	}

	var reqBody = {
		'username': username,
		'pwd': pwd
	}

	var dat = {
		'url': 'http://127.0.0.1:8001/user/' + username,
		'method': 'POST',
		'req_body': JSON.stringify(reqBody)
	};
	$.ajax({
		url  : "/api",
		type : 'post',
		data : JSON.stringify(dat),
		statusCode: {
			500: function() {
				callback(null, "Internal error");
			}
		},
		complete: function(xhr, textStatus) {
			if (xhr.status >= 400) {
				callback(null, "Error of Signin");
				return;
			}
		}
	}).done(function(data, statusText, xhr){
		if (xhr.status >= 400) {
			callback(null, "Error of Signin");
			return;
		}
		uname = username;

		callback(data, null);
	});
}

// done: 完成
function getUserId(callback) {
	if (uname === "") {
		return
	}
	var dat = {
		'url': 'http://127.0.0.1:8001/user/' + uname,
		'method': 'GET'
	};

	$.ajax({
		url: '/api',
		type: 'post',
		data: JSON.stringify(dat),
		headers: {'X-Session-Id': session},
		statusCode: {
			500: function() {
				callback(null, "Internal Error");
			}
		},
		complete: function(xhr, textStatus) {
			if (xhr.status >= 400) {
				callback(null, "Error of getUserId");
				return;
			}
		}
	}).done(function (data, statusText, xhr) {
		callback(data, null);
	});
}

// note: Video operations
function createVideo(vname, callback) {
	var reqBody = {
		'user_id': uid,
		'name': vname
	};

	var dat = {
		'url': 'http://127.0.0.1:8001/user/' + uname + '/videos',
		'method': 'POST',
		'req_body': JSON.stringify(reqBody)
	};

	$.ajax({
		url  : '/api',
		type : 'post',
		data : JSON.stringify(dat),
		headers: {'X-Session-Id': session},
		statusCode: {
			500: function() {
				callback(null, "Internal error");
			}
		},
		complete: function(xhr, textStatus) {
			if (xhr.status >= 400) {
				callback(null, "Error of Signin");
				return;
			}
		}
	}).done(function(data, statusText, xhr){
		if (xhr.status >= 400) {
			callback(null, "Error of Signin");
			return;
		}
		callback(data, null);
	});
}

// done: 完成
function listAllVideos(callback) {
  var dat = {
    'url': 'http://127.0.0.1:8001/user/' + uname + '/videos',
    'method': 'GET',
    'req_body': ''
  };

  $.ajax({
    url  : '/api',
    type : 'post',
    data : JSON.stringify(dat),
    headers: {'X-Session-Id': session},
    statusCode: {
      500: function() {
        callback(null, "Internal error");
      }
    },
    complete: function(xhr, textStatus) {
      if (xhr.status >= 400) {
        callback(null, "Error of Signin");
        return;
      }
    }
  }).done(function(data, statusText, xhr){
    if (xhr.status >= 400) {
      callback(null, "Error of Signin");
      return;
    }
    callback(data, null);
  });
}

function deleteVideo(vid, callback) {
  var dat = {
    'url': 'http://127.0.0.1:8001/user/' + uname + '/videos/' + vid,
    'method': 'DELETE',
    'req_body': ''
  };

  $.ajax({
    url  : '/api',
    type : 'post',
    data : JSON.stringify(dat),
    headers: {'X-Session-Id': session},
    statusCode: {
      500: function() {
        callback(null, "Internal error");
      }
    },
    complete: function(xhr, textStatus) {
      if (xhr.status >= 400) {
        callback(null, "Error of Signin");
        return;
      }
    }
  }).done(function(data, statusText, xhr){
    if (xhr.status >= 400) {
      callback(null, "Error of Signin");
      return;
    }
    callback(data, null);
  });
}

// note: Comments operations
function postComment(vid, content, callback) {
 var reqBody = {
  	'user_id': uid,
  	'content': content
  }

  var dat = {
    'url': 'http://127.0.0.1:8001/videos/' + vid + '/comments',
    'method': 'POST',
    'req_body': JSON.stringify(reqBody)
  };

  $.ajax({
    url  : '/api',
    type : 'post',
    data : JSON.stringify(dat),
    headers: {'X-Session-Id': session},
    statusCode: {
      500: function() {
        callback(null, "Internal error");
      }
    },
    complete: function(xhr, textStatus) {
      if (xhr.status >= 400) {
        callback(null, "Error of Signin");
        return;
      }
    }
  }).done(function(data, statusText, xhr){
    if (xhr.status >= 400) {
      callback(null, "Error of Signin");
      return;
    }
    callback(data, null);
  });
}

function listAllComments(vid, callback) {
  var dat = {
    'url': 'http://127.0.0.1:8001/videos/' + vid + '/comments',
    'method': 'GET',
    'req_body': ''
  };

  $.ajax({
    url  : '/api',
    type : 'post',
    data : JSON.stringify(dat),
    headers: {'X-Session-Id': session},
    statusCode: {
      500: function() {
        callback(null, "Internal error");
      }
    },
    complete: function(xhr, textStatus) {
      if (xhr.status >= 400) {
        callback(null, "Error of Signin");
        return;
      }
    }
  }).done(function(data, statusText, xhr){
    if (xhr.status >= 400) {
      callback(null, "Error of Signin");
      return;
    }
    callback(data, null);
  });
}
