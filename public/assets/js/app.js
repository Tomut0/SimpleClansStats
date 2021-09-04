/*
 * ATTENTION: An "eval-source-map" devtool has been used.
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file with attached SourceMaps in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
/******/ (() => { // webpackBootstrap
/******/ 	var __webpack_modules__ = ({

/***/ "./resources/js/detail_modal.js":
/*!**************************************!*\
  !*** ./resources/js/detail_modal.js ***!
  \**************************************/
/***/ (() => {

eval("function _createForOfIteratorHelper(o, allowArrayLike) { var it = typeof Symbol !== \"undefined\" && o[Symbol.iterator] || o[\"@@iterator\"]; if (!it) { if (Array.isArray(o) || (it = _unsupportedIterableToArray(o)) || allowArrayLike && o && typeof o.length === \"number\") { if (it) o = it; var i = 0; var F = function F() {}; return { s: F, n: function n() { if (i >= o.length) return { done: true }; return { done: false, value: o[i++] }; }, e: function e(_e) { throw _e; }, f: F }; } throw new TypeError(\"Invalid attempt to iterate non-iterable instance.\\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.\"); } var normalCompletion = true, didErr = false, err; return { s: function s() { it = it.call(o); }, n: function n() { var step = it.next(); normalCompletion = step.done; return step; }, e: function e(_e2) { didErr = true; err = _e2; }, f: function f() { try { if (!normalCompletion && it[\"return\"] != null) it[\"return\"](); } finally { if (didErr) throw err; } } }; }\n\nfunction _unsupportedIterableToArray(o, minLen) { if (!o) return; if (typeof o === \"string\") return _arrayLikeToArray(o, minLen); var n = Object.prototype.toString.call(o).slice(8, -1); if (n === \"Object\" && o.constructor) n = o.constructor.name; if (n === \"Map\" || n === \"Set\") return Array.from(o); if (n === \"Arguments\" || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)) return _arrayLikeToArray(o, minLen); }\n\nfunction _arrayLikeToArray(arr, len) { if (len == null || len > arr.length) len = arr.length; for (var i = 0, arr2 = new Array(len); i < len; i++) { arr2[i] = arr[i]; } return arr2; }\n\nvar baseClanUrl = \"/detail/clan/\";\nvar basePlayerUrl = \"/detail/player/\";\n\nfunction setupModalHolder() {\n  var modalHolder = document.getElementById(\"modalHolder\");\n\n  if (!modalHolder) {\n    modalHolder = document.createElement(\"div\");\n    modalHolder.id = \"modalHolder\";\n    document.body.appendChild(modalHolder);\n  }\n}\n\nfunction getModal() {\n  var element = document.querySelector('#detailModal');\n\n  if (!element) {\n    throw new Error(\"Modal not loaded\");\n  }\n\n  return new bootstrap.Modal(element);\n}\n\nfunction removeBackdrop() {\n  var backdrop = document.body.querySelector(\".modal-backdrop.fade.show\");\n\n  if (backdrop) {\n    document.body.removeChild(backdrop);\n  }\n}\n\nfunction showModal(url, id) {\n  if (!id) {\n    console.log(\"invalid id\");\n    return;\n  } //when opening a modal from another, the backdrop doesn't disappear, so this takes care of it\n\n\n  removeBackdrop();\n  setupModalHolder();\n  fetch(url + id).then(function (response) {\n    return response.text();\n  }).then(function (text) {\n    document.querySelector('#modalHolder').innerHTML = text;\n    getModal().show(); //registering again, in case the modal has modal-openers\n\n    registerModalOpeners();\n  });\n}\n\nfunction showPlayer(name) {\n  showModal(basePlayerUrl, name);\n}\n\nfunction showClan(tag) {\n  showModal(baseClanUrl, tag);\n}\n\nfunction registerModalOpeners() {\n  var elements = document.querySelectorAll(\".modal-opener\");\n\n  var _iterator = _createForOfIteratorHelper(elements),\n      _step;\n\n  try {\n    var _loop = function _loop() {\n      var element = _step.value;\n      var tag = element.dataset.tag;\n      var nick = element.dataset.nick;\n\n      if (tag) {\n        element.onclick = function () {\n          return showClan(tag);\n        };\n      } else if (nick) {\n        element.onclick = function () {\n          return showPlayer(nick);\n        };\n      }\n    };\n\n    for (_iterator.s(); !(_step = _iterator.n()).done;) {\n      _loop();\n    }\n  } catch (err) {\n    _iterator.e(err);\n  } finally {\n    _iterator.f();\n  }\n}\n\nregisterModalOpeners();//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vLi9yZXNvdXJjZXMvanMvZGV0YWlsX21vZGFsLmpzPzRkOWMiXSwibmFtZXMiOlsiYmFzZUNsYW5VcmwiLCJiYXNlUGxheWVyVXJsIiwic2V0dXBNb2RhbEhvbGRlciIsIm1vZGFsSG9sZGVyIiwiZG9jdW1lbnQiLCJnZXRFbGVtZW50QnlJZCIsImNyZWF0ZUVsZW1lbnQiLCJpZCIsImJvZHkiLCJhcHBlbmRDaGlsZCIsImdldE1vZGFsIiwiZWxlbWVudCIsInF1ZXJ5U2VsZWN0b3IiLCJFcnJvciIsImJvb3RzdHJhcCIsIk1vZGFsIiwicmVtb3ZlQmFja2Ryb3AiLCJiYWNrZHJvcCIsInJlbW92ZUNoaWxkIiwic2hvd01vZGFsIiwidXJsIiwiY29uc29sZSIsImxvZyIsImZldGNoIiwidGhlbiIsInJlc3BvbnNlIiwidGV4dCIsImlubmVySFRNTCIsInNob3ciLCJyZWdpc3Rlck1vZGFsT3BlbmVycyIsInNob3dQbGF5ZXIiLCJuYW1lIiwic2hvd0NsYW4iLCJ0YWciLCJlbGVtZW50cyIsInF1ZXJ5U2VsZWN0b3JBbGwiLCJkYXRhc2V0IiwibmljayIsIm9uY2xpY2siXSwibWFwcGluZ3MiOiI7Ozs7OztBQUFBLElBQU1BLFdBQVcsR0FBRyxlQUFwQjtBQUNBLElBQU1DLGFBQWEsR0FBRyxpQkFBdEI7O0FBRUEsU0FBU0MsZ0JBQVQsR0FBNEI7QUFDeEIsTUFBSUMsV0FBVyxHQUFHQyxRQUFRLENBQUNDLGNBQVQsQ0FBd0IsYUFBeEIsQ0FBbEI7O0FBQ0EsTUFBSSxDQUFDRixXQUFMLEVBQWtCO0FBQ2RBLElBQUFBLFdBQVcsR0FBR0MsUUFBUSxDQUFDRSxhQUFULENBQXVCLEtBQXZCLENBQWQ7QUFDQUgsSUFBQUEsV0FBVyxDQUFDSSxFQUFaLEdBQWlCLGFBQWpCO0FBQ0FILElBQUFBLFFBQVEsQ0FBQ0ksSUFBVCxDQUFjQyxXQUFkLENBQTBCTixXQUExQjtBQUNIO0FBQ0o7O0FBRUQsU0FBU08sUUFBVCxHQUFvQjtBQUNoQixNQUFJQyxPQUFPLEdBQUdQLFFBQVEsQ0FBQ1EsYUFBVCxDQUF1QixjQUF2QixDQUFkOztBQUNBLE1BQUksQ0FBQ0QsT0FBTCxFQUFjO0FBQ1YsVUFBTSxJQUFJRSxLQUFKLENBQVUsa0JBQVYsQ0FBTjtBQUNIOztBQUNELFNBQU8sSUFBSUMsU0FBUyxDQUFDQyxLQUFkLENBQW9CSixPQUFwQixDQUFQO0FBQ0g7O0FBRUQsU0FBU0ssY0FBVCxHQUEwQjtBQUN0QixNQUFJQyxRQUFRLEdBQUdiLFFBQVEsQ0FBQ0ksSUFBVCxDQUFjSSxhQUFkLENBQTRCLDJCQUE1QixDQUFmOztBQUNBLE1BQUlLLFFBQUosRUFBYztBQUNWYixJQUFBQSxRQUFRLENBQUNJLElBQVQsQ0FBY1UsV0FBZCxDQUEwQkQsUUFBMUI7QUFDSDtBQUNKOztBQUVELFNBQVNFLFNBQVQsQ0FBbUJDLEdBQW5CLEVBQXdCYixFQUF4QixFQUE0QjtBQUN4QixNQUFJLENBQUNBLEVBQUwsRUFBUztBQUNMYyxJQUFBQSxPQUFPLENBQUNDLEdBQVIsQ0FBWSxZQUFaO0FBQ0E7QUFDSCxHQUp1QixDQUt4Qjs7O0FBQ0FOLEVBQUFBLGNBQWM7QUFDZGQsRUFBQUEsZ0JBQWdCO0FBQ2hCcUIsRUFBQUEsS0FBSyxDQUFDSCxHQUFHLEdBQUdiLEVBQVAsQ0FBTCxDQUFnQmlCLElBQWhCLENBQXFCLFVBQVVDLFFBQVYsRUFBb0I7QUFDckMsV0FBT0EsUUFBUSxDQUFDQyxJQUFULEVBQVA7QUFDSCxHQUZELEVBRUdGLElBRkgsQ0FFUSxVQUFVRSxJQUFWLEVBQWdCO0FBQ3BCdEIsSUFBQUEsUUFBUSxDQUFDUSxhQUFULENBQXVCLGNBQXZCLEVBQXVDZSxTQUF2QyxHQUFtREQsSUFBbkQ7QUFDQWhCLElBQUFBLFFBQVEsR0FBR2tCLElBQVgsR0FGb0IsQ0FHcEI7O0FBQ0FDLElBQUFBLG9CQUFvQjtBQUN2QixHQVBEO0FBUUg7O0FBRUQsU0FBU0MsVUFBVCxDQUFvQkMsSUFBcEIsRUFBMEI7QUFDdEJaLEVBQUFBLFNBQVMsQ0FBQ2xCLGFBQUQsRUFBZ0I4QixJQUFoQixDQUFUO0FBQ0g7O0FBRUQsU0FBU0MsUUFBVCxDQUFrQkMsR0FBbEIsRUFBdUI7QUFDbkJkLEVBQUFBLFNBQVMsQ0FBQ25CLFdBQUQsRUFBY2lDLEdBQWQsQ0FBVDtBQUNIOztBQUVELFNBQVNKLG9CQUFULEdBQWdDO0FBQzVCLE1BQUlLLFFBQVEsR0FBRzlCLFFBQVEsQ0FBQytCLGdCQUFULENBQTBCLGVBQTFCLENBQWY7O0FBRDRCLDZDQUVSRCxRQUZRO0FBQUE7O0FBQUE7QUFBQTtBQUFBLFVBRW5CdkIsT0FGbUI7QUFHeEIsVUFBSXNCLEdBQUcsR0FBR3RCLE9BQU8sQ0FBQ3lCLE9BQVIsQ0FBZ0JILEdBQTFCO0FBQ0EsVUFBSUksSUFBSSxHQUFHMUIsT0FBTyxDQUFDeUIsT0FBUixDQUFnQkMsSUFBM0I7O0FBQ0EsVUFBSUosR0FBSixFQUFTO0FBQ0x0QixRQUFBQSxPQUFPLENBQUMyQixPQUFSLEdBQWtCO0FBQUEsaUJBQU1OLFFBQVEsQ0FBQ0MsR0FBRCxDQUFkO0FBQUEsU0FBbEI7QUFDSCxPQUZELE1BRU8sSUFBSUksSUFBSixFQUFVO0FBQ2IxQixRQUFBQSxPQUFPLENBQUMyQixPQUFSLEdBQWtCO0FBQUEsaUJBQU1SLFVBQVUsQ0FBQ08sSUFBRCxDQUFoQjtBQUFBLFNBQWxCO0FBQ0g7QUFUdUI7O0FBRTVCLHdEQUE4QjtBQUFBO0FBUTdCO0FBVjJCO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFXL0I7O0FBRURSLG9CQUFvQiIsInNvdXJjZXNDb250ZW50IjpbImNvbnN0IGJhc2VDbGFuVXJsID0gXCIvZGV0YWlsL2NsYW4vXCI7XHJcbmNvbnN0IGJhc2VQbGF5ZXJVcmwgPSBcIi9kZXRhaWwvcGxheWVyL1wiO1xyXG5cclxuZnVuY3Rpb24gc2V0dXBNb2RhbEhvbGRlcigpIHtcclxuICAgIGxldCBtb2RhbEhvbGRlciA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKFwibW9kYWxIb2xkZXJcIik7XHJcbiAgICBpZiAoIW1vZGFsSG9sZGVyKSB7XHJcbiAgICAgICAgbW9kYWxIb2xkZXIgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwiZGl2XCIpO1xyXG4gICAgICAgIG1vZGFsSG9sZGVyLmlkID0gXCJtb2RhbEhvbGRlclwiO1xyXG4gICAgICAgIGRvY3VtZW50LmJvZHkuYXBwZW5kQ2hpbGQobW9kYWxIb2xkZXIpO1xyXG4gICAgfVxyXG59XHJcblxyXG5mdW5jdGlvbiBnZXRNb2RhbCgpIHtcclxuICAgIGxldCBlbGVtZW50ID0gZG9jdW1lbnQucXVlcnlTZWxlY3RvcignI2RldGFpbE1vZGFsJyk7XHJcbiAgICBpZiAoIWVsZW1lbnQpIHtcclxuICAgICAgICB0aHJvdyBuZXcgRXJyb3IoXCJNb2RhbCBub3QgbG9hZGVkXCIpO1xyXG4gICAgfVxyXG4gICAgcmV0dXJuIG5ldyBib290c3RyYXAuTW9kYWwoZWxlbWVudCk7XHJcbn1cclxuXHJcbmZ1bmN0aW9uIHJlbW92ZUJhY2tkcm9wKCkge1xyXG4gICAgbGV0IGJhY2tkcm9wID0gZG9jdW1lbnQuYm9keS5xdWVyeVNlbGVjdG9yKFwiLm1vZGFsLWJhY2tkcm9wLmZhZGUuc2hvd1wiKTtcclxuICAgIGlmIChiYWNrZHJvcCkge1xyXG4gICAgICAgIGRvY3VtZW50LmJvZHkucmVtb3ZlQ2hpbGQoYmFja2Ryb3ApO1xyXG4gICAgfVxyXG59XHJcblxyXG5mdW5jdGlvbiBzaG93TW9kYWwodXJsLCBpZCkge1xyXG4gICAgaWYgKCFpZCkge1xyXG4gICAgICAgIGNvbnNvbGUubG9nKFwiaW52YWxpZCBpZFwiKVxyXG4gICAgICAgIHJldHVybjtcclxuICAgIH1cclxuICAgIC8vd2hlbiBvcGVuaW5nIGEgbW9kYWwgZnJvbSBhbm90aGVyLCB0aGUgYmFja2Ryb3AgZG9lc24ndCBkaXNhcHBlYXIsIHNvIHRoaXMgdGFrZXMgY2FyZSBvZiBpdFxyXG4gICAgcmVtb3ZlQmFja2Ryb3AoKTtcclxuICAgIHNldHVwTW9kYWxIb2xkZXIoKTtcclxuICAgIGZldGNoKHVybCArIGlkKS50aGVuKGZ1bmN0aW9uIChyZXNwb25zZSkge1xyXG4gICAgICAgIHJldHVybiByZXNwb25zZS50ZXh0KCk7XHJcbiAgICB9KS50aGVuKGZ1bmN0aW9uICh0ZXh0KSB7XHJcbiAgICAgICAgZG9jdW1lbnQucXVlcnlTZWxlY3RvcignI21vZGFsSG9sZGVyJykuaW5uZXJIVE1MID0gdGV4dDtcclxuICAgICAgICBnZXRNb2RhbCgpLnNob3coKTtcclxuICAgICAgICAvL3JlZ2lzdGVyaW5nIGFnYWluLCBpbiBjYXNlIHRoZSBtb2RhbCBoYXMgbW9kYWwtb3BlbmVyc1xyXG4gICAgICAgIHJlZ2lzdGVyTW9kYWxPcGVuZXJzKCk7XHJcbiAgICB9KTtcclxufVxyXG5cclxuZnVuY3Rpb24gc2hvd1BsYXllcihuYW1lKSB7XHJcbiAgICBzaG93TW9kYWwoYmFzZVBsYXllclVybCwgbmFtZSk7XHJcbn1cclxuXHJcbmZ1bmN0aW9uIHNob3dDbGFuKHRhZykge1xyXG4gICAgc2hvd01vZGFsKGJhc2VDbGFuVXJsLCB0YWcpO1xyXG59XHJcblxyXG5mdW5jdGlvbiByZWdpc3Rlck1vZGFsT3BlbmVycygpIHtcclxuICAgIGxldCBlbGVtZW50cyA9IGRvY3VtZW50LnF1ZXJ5U2VsZWN0b3JBbGwoXCIubW9kYWwtb3BlbmVyXCIpO1xyXG4gICAgZm9yIChsZXQgZWxlbWVudCBvZiBlbGVtZW50cykge1xyXG4gICAgICAgIGxldCB0YWcgPSBlbGVtZW50LmRhdGFzZXQudGFnO1xyXG4gICAgICAgIGxldCBuaWNrID0gZWxlbWVudC5kYXRhc2V0Lm5pY2s7XHJcbiAgICAgICAgaWYgKHRhZykge1xyXG4gICAgICAgICAgICBlbGVtZW50Lm9uY2xpY2sgPSAoKSA9PiBzaG93Q2xhbih0YWcpXHJcbiAgICAgICAgfSBlbHNlIGlmIChuaWNrKSB7XHJcbiAgICAgICAgICAgIGVsZW1lbnQub25jbGljayA9ICgpID0+IHNob3dQbGF5ZXIobmljayk7XHJcbiAgICAgICAgfVxyXG4gICAgfVxyXG59XHJcblxyXG5yZWdpc3Rlck1vZGFsT3BlbmVycygpO1xyXG4iXSwiZmlsZSI6Ii4vcmVzb3VyY2VzL2pzL2RldGFpbF9tb2RhbC5qcy5qcyIsInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///./resources/js/detail_modal.js\n");

/***/ }),

/***/ "./resources/js/filter.js":
/*!********************************!*\
  !*** ./resources/js/filter.js ***!
  \********************************/
/***/ (() => {

eval("function _createForOfIteratorHelper(o, allowArrayLike) { var it = typeof Symbol !== \"undefined\" && o[Symbol.iterator] || o[\"@@iterator\"]; if (!it) { if (Array.isArray(o) || (it = _unsupportedIterableToArray(o)) || allowArrayLike && o && typeof o.length === \"number\") { if (it) o = it; var i = 0; var F = function F() {}; return { s: F, n: function n() { if (i >= o.length) return { done: true }; return { done: false, value: o[i++] }; }, e: function e(_e) { throw _e; }, f: F }; } throw new TypeError(\"Invalid attempt to iterate non-iterable instance.\\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.\"); } var normalCompletion = true, didErr = false, err; return { s: function s() { it = it.call(o); }, n: function n() { var step = it.next(); normalCompletion = step.done; return step; }, e: function e(_e2) { didErr = true; err = _e2; }, f: function f() { try { if (!normalCompletion && it[\"return\"] != null) it[\"return\"](); } finally { if (didErr) throw err; } } }; }\n\nfunction _unsupportedIterableToArray(o, minLen) { if (!o) return; if (typeof o === \"string\") return _arrayLikeToArray(o, minLen); var n = Object.prototype.toString.call(o).slice(8, -1); if (n === \"Object\" && o.constructor) n = o.constructor.name; if (n === \"Map\" || n === \"Set\") return Array.from(o); if (n === \"Arguments\" || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)) return _arrayLikeToArray(o, minLen); }\n\nfunction _arrayLikeToArray(arr, len) { if (len == null || len > arr.length) len = arr.length; for (var i = 0, arr2 = new Array(len); i < len; i++) { arr2[i] = arr[i]; } return arr2; }\n\nfunction filter(rows, text) {\n  var _iterator = _createForOfIteratorHelper(rows),\n      _step;\n\n  try {\n    outter: for (_iterator.s(); !(_step = _iterator.n()).done;) {\n      var row = _step.value;\n\n      var _iterator2 = _createForOfIteratorHelper(row.children),\n          _step2;\n\n      try {\n        for (_iterator2.s(); !(_step2 = _iterator2.n()).done;) {\n          var column = _step2.value;\n\n          if (column.innerText.toLowerCase().includes(text.toLowerCase())) {\n            row.hidden = false;\n            continue outter;\n          }\n        }\n      } catch (err) {\n        _iterator2.e(err);\n      } finally {\n        _iterator2.f();\n      }\n\n      row.hidden = true;\n    }\n  } catch (err) {\n    _iterator.e(err);\n  } finally {\n    _iterator.f();\n  }\n}\n\nvar filterInput = document.getElementById(\"filter\");\n\nvar listener = function listener() {\n  var text = filterInput.value;\n  var rows = document.getElementsByClassName(\"data_row\");\n  filter(rows, text);\n};\n\nfilterInput.addEventListener(\"input\", listener);\nwindow.addEventListener(\"load\", listener);//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vLi9yZXNvdXJjZXMvanMvZmlsdGVyLmpzPzc5N2YiXSwibmFtZXMiOlsiZmlsdGVyIiwicm93cyIsInRleHQiLCJvdXR0ZXIiLCJyb3ciLCJjaGlsZHJlbiIsImNvbHVtbiIsImlubmVyVGV4dCIsInRvTG93ZXJDYXNlIiwiaW5jbHVkZXMiLCJoaWRkZW4iLCJmaWx0ZXJJbnB1dCIsImRvY3VtZW50IiwiZ2V0RWxlbWVudEJ5SWQiLCJsaXN0ZW5lciIsInZhbHVlIiwiZ2V0RWxlbWVudHNCeUNsYXNzTmFtZSIsImFkZEV2ZW50TGlzdGVuZXIiLCJ3aW5kb3ciXSwibWFwcGluZ3MiOiI7Ozs7OztBQUFBLFNBQVNBLE1BQVQsQ0FBZ0JDLElBQWhCLEVBQXNCQyxJQUF0QixFQUE0QjtBQUFBLDZDQUVGRCxJQUZFO0FBQUE7O0FBQUE7QUFDeEJFLElBQUFBLE1BRHdCLEVBRXBCLG9EQUF3QjtBQUFBLFVBQWJDLEdBQWE7O0FBQUEsa0RBQ0NBLEdBQUcsQ0FBQ0MsUUFETDtBQUFBOztBQUFBO0FBQ3BCLCtEQUFtQztBQUFBLGNBQXhCQyxNQUF3Qjs7QUFDL0IsY0FBSUEsTUFBTSxDQUFDQyxTQUFQLENBQWlCQyxXQUFqQixHQUErQkMsUUFBL0IsQ0FBd0NQLElBQUksQ0FBQ00sV0FBTCxFQUF4QyxDQUFKLEVBQWlFO0FBQzdESixZQUFBQSxHQUFHLENBQUNNLE1BQUosR0FBYSxLQUFiO0FBQ0EscUJBQVNQLE1BQVQ7QUFDSDtBQUNKO0FBTm1CO0FBQUE7QUFBQTtBQUFBO0FBQUE7O0FBT3BCQyxNQUFBQSxHQUFHLENBQUNNLE1BQUosR0FBYSxJQUFiO0FBQ0g7QUFWbUI7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQVczQjs7QUFFRCxJQUFNQyxXQUFXLEdBQUdDLFFBQVEsQ0FBQ0MsY0FBVCxDQUF3QixRQUF4QixDQUFwQjs7QUFDQSxJQUFNQyxRQUFRLEdBQUcsU0FBWEEsUUFBVyxHQUFZO0FBQ3pCLE1BQUlaLElBQUksR0FBR1MsV0FBVyxDQUFDSSxLQUF2QjtBQUNBLE1BQU1kLElBQUksR0FBR1csUUFBUSxDQUFDSSxzQkFBVCxDQUFnQyxVQUFoQyxDQUFiO0FBQ0FoQixFQUFBQSxNQUFNLENBQUNDLElBQUQsRUFBT0MsSUFBUCxDQUFOO0FBQ0gsQ0FKRDs7QUFLQVMsV0FBVyxDQUFDTSxnQkFBWixDQUE2QixPQUE3QixFQUFzQ0gsUUFBdEM7QUFDQUksTUFBTSxDQUFDRCxnQkFBUCxDQUF3QixNQUF4QixFQUFnQ0gsUUFBaEMiLCJzb3VyY2VzQ29udGVudCI6WyJmdW5jdGlvbiBmaWx0ZXIocm93cywgdGV4dCkge1xyXG4gICAgb3V0dGVyOlxyXG4gICAgICAgIGZvciAoY29uc3Qgcm93IG9mIHJvd3MpIHtcclxuICAgICAgICAgICAgZm9yIChjb25zdCBjb2x1bW4gb2Ygcm93LmNoaWxkcmVuKSB7XHJcbiAgICAgICAgICAgICAgICBpZiAoY29sdW1uLmlubmVyVGV4dC50b0xvd2VyQ2FzZSgpLmluY2x1ZGVzKHRleHQudG9Mb3dlckNhc2UoKSkpIHtcclxuICAgICAgICAgICAgICAgICAgICByb3cuaGlkZGVuID0gZmFsc2U7XHJcbiAgICAgICAgICAgICAgICAgICAgY29udGludWUgb3V0dGVyO1xyXG4gICAgICAgICAgICAgICAgfVxyXG4gICAgICAgICAgICB9XHJcbiAgICAgICAgICAgIHJvdy5oaWRkZW4gPSB0cnVlO1xyXG4gICAgICAgIH1cclxufVxyXG5cclxuY29uc3QgZmlsdGVySW5wdXQgPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcImZpbHRlclwiKTtcclxuY29uc3QgbGlzdGVuZXIgPSBmdW5jdGlvbiAoKSB7XHJcbiAgICBsZXQgdGV4dCA9IGZpbHRlcklucHV0LnZhbHVlO1xyXG4gICAgY29uc3Qgcm93cyA9IGRvY3VtZW50LmdldEVsZW1lbnRzQnlDbGFzc05hbWUoXCJkYXRhX3Jvd1wiKTtcclxuICAgIGZpbHRlcihyb3dzLCB0ZXh0KTtcclxufVxyXG5maWx0ZXJJbnB1dC5hZGRFdmVudExpc3RlbmVyKFwiaW5wdXRcIiwgbGlzdGVuZXIpO1xyXG53aW5kb3cuYWRkRXZlbnRMaXN0ZW5lcihcImxvYWRcIiwgbGlzdGVuZXIpO1xyXG4iXSwiZmlsZSI6Ii4vcmVzb3VyY2VzL2pzL2ZpbHRlci5qcy5qcyIsInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///./resources/js/filter.js\n");

/***/ }),

/***/ "./resources/js/trim_names.js":
/*!************************************!*\
  !*** ./resources/js/trim_names.js ***!
  \************************************/
/***/ (() => {

eval("function _createForOfIteratorHelper(o, allowArrayLike) { var it = typeof Symbol !== \"undefined\" && o[Symbol.iterator] || o[\"@@iterator\"]; if (!it) { if (Array.isArray(o) || (it = _unsupportedIterableToArray(o)) || allowArrayLike && o && typeof o.length === \"number\") { if (it) o = it; var i = 0; var F = function F() {}; return { s: F, n: function n() { if (i >= o.length) return { done: true }; return { done: false, value: o[i++] }; }, e: function e(_e) { throw _e; }, f: F }; } throw new TypeError(\"Invalid attempt to iterate non-iterable instance.\\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.\"); } var normalCompletion = true, didErr = false, err; return { s: function s() { it = it.call(o); }, n: function n() { var step = it.next(); normalCompletion = step.done; return step; }, e: function e(_e2) { didErr = true; err = _e2; }, f: function f() { try { if (!normalCompletion && it[\"return\"] != null) it[\"return\"](); } finally { if (didErr) throw err; } } }; }\n\nfunction _unsupportedIterableToArray(o, minLen) { if (!o) return; if (typeof o === \"string\") return _arrayLikeToArray(o, minLen); var n = Object.prototype.toString.call(o).slice(8, -1); if (n === \"Object\" && o.constructor) n = o.constructor.name; if (n === \"Map\" || n === \"Set\") return Array.from(o); if (n === \"Arguments\" || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)) return _arrayLikeToArray(o, minLen); }\n\nfunction _arrayLikeToArray(arr, len) { if (len == null || len > arr.length) len = arr.length; for (var i = 0, arr2 = new Array(len); i < len; i++) { arr2[i] = arr[i]; } return arr2; }\n\nfunction trimNames(className, maxLength) {\n  var split = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : false;\n  var elements = document.getElementsByClassName(className);\n\n  var _iterator = _createForOfIteratorHelper(elements),\n      _step;\n\n  try {\n    for (_iterator.s(); !(_step = _iterator.n()).done;) {\n      var element = _step.value;\n      var text = element.innerText;\n\n      if (text.length > maxLength) {\n        if (split) {\n          text = text.split(\" \")[1];\n        }\n\n        var original = text;\n        element.title = original;\n        text = text.substr(0, maxLength - 3) + \"…\";\n        element.innerHTML = element.innerHTML.replace(original, text);\n      }\n    }\n  } catch (err) {\n    _iterator.e(err);\n  } finally {\n    _iterator.f();\n  }\n}\n\ntrimNames(\"player_name\", 15, true);\ntrimNames(\"clan_name\", 20);//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vLi9yZXNvdXJjZXMvanMvdHJpbV9uYW1lcy5qcz9kYzg3Il0sIm5hbWVzIjpbInRyaW1OYW1lcyIsImNsYXNzTmFtZSIsIm1heExlbmd0aCIsInNwbGl0IiwiZWxlbWVudHMiLCJkb2N1bWVudCIsImdldEVsZW1lbnRzQnlDbGFzc05hbWUiLCJlbGVtZW50IiwidGV4dCIsImlubmVyVGV4dCIsImxlbmd0aCIsIm9yaWdpbmFsIiwidGl0bGUiLCJzdWJzdHIiLCJpbm5lckhUTUwiLCJyZXBsYWNlIl0sIm1hcHBpbmdzIjoiOzs7Ozs7QUFBQSxTQUFTQSxTQUFULENBQW1CQyxTQUFuQixFQUE4QkMsU0FBOUIsRUFBd0Q7QUFBQSxNQUFmQyxLQUFlLHVFQUFQLEtBQU87QUFDcEQsTUFBTUMsUUFBUSxHQUFHQyxRQUFRLENBQUNDLHNCQUFULENBQWdDTCxTQUFoQyxDQUFqQjs7QUFEb0QsNkNBRTlCRyxRQUY4QjtBQUFBOztBQUFBO0FBRXBELHdEQUFnQztBQUFBLFVBQXJCRyxPQUFxQjtBQUM1QixVQUFJQyxJQUFJLEdBQUdELE9BQU8sQ0FBQ0UsU0FBbkI7O0FBQ0EsVUFBSUQsSUFBSSxDQUFDRSxNQUFMLEdBQWNSLFNBQWxCLEVBQTZCO0FBQ3pCLFlBQUlDLEtBQUosRUFBVztBQUNQSyxVQUFBQSxJQUFJLEdBQUdBLElBQUksQ0FBQ0wsS0FBTCxDQUFXLEdBQVgsRUFBZ0IsQ0FBaEIsQ0FBUDtBQUNIOztBQUNELFlBQUlRLFFBQVEsR0FBR0gsSUFBZjtBQUNBRCxRQUFBQSxPQUFPLENBQUNLLEtBQVIsR0FBZ0JELFFBQWhCO0FBQ0FILFFBQUFBLElBQUksR0FBR0EsSUFBSSxDQUFDSyxNQUFMLENBQVksQ0FBWixFQUFlWCxTQUFTLEdBQUcsQ0FBM0IsSUFBZ0MsR0FBdkM7QUFDQUssUUFBQUEsT0FBTyxDQUFDTyxTQUFSLEdBQW9CUCxPQUFPLENBQUNPLFNBQVIsQ0FBa0JDLE9BQWxCLENBQTBCSixRQUExQixFQUFvQ0gsSUFBcEMsQ0FBcEI7QUFDSDtBQUNKO0FBYm1EO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFjdkQ7O0FBQ0RSLFNBQVMsQ0FBQyxhQUFELEVBQWdCLEVBQWhCLEVBQW9CLElBQXBCLENBQVQ7QUFDQUEsU0FBUyxDQUFDLFdBQUQsRUFBYyxFQUFkLENBQVQiLCJzb3VyY2VzQ29udGVudCI6WyJmdW5jdGlvbiB0cmltTmFtZXMoY2xhc3NOYW1lLCBtYXhMZW5ndGgsIHNwbGl0ID0gZmFsc2UpIHtcclxuICAgIGNvbnN0IGVsZW1lbnRzID0gZG9jdW1lbnQuZ2V0RWxlbWVudHNCeUNsYXNzTmFtZShjbGFzc05hbWUpO1xyXG4gICAgZm9yIChjb25zdCBlbGVtZW50IG9mIGVsZW1lbnRzKSB7XHJcbiAgICAgICAgbGV0IHRleHQgPSBlbGVtZW50LmlubmVyVGV4dDtcclxuICAgICAgICBpZiAodGV4dC5sZW5ndGggPiBtYXhMZW5ndGgpIHtcclxuICAgICAgICAgICAgaWYgKHNwbGl0KSB7XHJcbiAgICAgICAgICAgICAgICB0ZXh0ID0gdGV4dC5zcGxpdChcIiBcIilbMV07XHJcbiAgICAgICAgICAgIH1cclxuICAgICAgICAgICAgbGV0IG9yaWdpbmFsID0gdGV4dDtcclxuICAgICAgICAgICAgZWxlbWVudC50aXRsZSA9IG9yaWdpbmFsO1xyXG4gICAgICAgICAgICB0ZXh0ID0gdGV4dC5zdWJzdHIoMCwgbWF4TGVuZ3RoIC0gMykgKyBcIuKAplwiO1xyXG4gICAgICAgICAgICBlbGVtZW50LmlubmVySFRNTCA9IGVsZW1lbnQuaW5uZXJIVE1MLnJlcGxhY2Uob3JpZ2luYWwsIHRleHQpO1xyXG4gICAgICAgIH1cclxuICAgIH1cclxufVxyXG50cmltTmFtZXMoXCJwbGF5ZXJfbmFtZVwiLCAxNSwgdHJ1ZSk7XHJcbnRyaW1OYW1lcyhcImNsYW5fbmFtZVwiLCAyMCk7XHJcbiJdLCJmaWxlIjoiLi9yZXNvdXJjZXMvanMvdHJpbV9uYW1lcy5qcy5qcyIsInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///./resources/js/trim_names.js\n");

/***/ }),

/***/ "./resources/sass/style.scss":
/*!***********************************!*\
  !*** ./resources/sass/style.scss ***!
  \***********************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n// extracted by mini-css-extract-plugin\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9yZXNvdXJjZXMvc2Fzcy9zdHlsZS5zY3NzLmpzIiwibWFwcGluZ3MiOiI7QUFBQSIsInNvdXJjZXMiOlsid2VicGFjazovLy8uL3Jlc291cmNlcy9zYXNzL3N0eWxlLnNjc3M/Y2VmNyJdLCJzb3VyY2VzQ29udGVudCI6WyIvLyBleHRyYWN0ZWQgYnkgbWluaS1jc3MtZXh0cmFjdC1wbHVnaW5cbmV4cG9ydCB7fTsiXSwibmFtZXMiOltdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///./resources/sass/style.scss\n");

/***/ })

/******/ 	});
/************************************************************************/
/******/ 	// The module cache
/******/ 	var __webpack_module_cache__ = {};
/******/ 	
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/ 		// Check if module is in cache
/******/ 		var cachedModule = __webpack_module_cache__[moduleId];
/******/ 		if (cachedModule !== undefined) {
/******/ 			return cachedModule.exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = __webpack_module_cache__[moduleId] = {
/******/ 			// no module.id needed
/******/ 			// no module.loaded needed
/******/ 			exports: {}
/******/ 		};
/******/ 	
/******/ 		// Execute the module function
/******/ 		__webpack_modules__[moduleId](module, module.exports, __webpack_require__);
/******/ 	
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/ 	
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = __webpack_modules__;
/******/ 	
/************************************************************************/
/******/ 	/* webpack/runtime/chunk loaded */
/******/ 	(() => {
/******/ 		var deferred = [];
/******/ 		__webpack_require__.O = (result, chunkIds, fn, priority) => {
/******/ 			if(chunkIds) {
/******/ 				priority = priority || 0;
/******/ 				for(var i = deferred.length; i > 0 && deferred[i - 1][2] > priority; i--) deferred[i] = deferred[i - 1];
/******/ 				deferred[i] = [chunkIds, fn, priority];
/******/ 				return;
/******/ 			}
/******/ 			var notFulfilled = Infinity;
/******/ 			for (var i = 0; i < deferred.length; i++) {
/******/ 				var [chunkIds, fn, priority] = deferred[i];
/******/ 				var fulfilled = true;
/******/ 				for (var j = 0; j < chunkIds.length; j++) {
/******/ 					if ((priority & 1 === 0 || notFulfilled >= priority) && Object.keys(__webpack_require__.O).every((key) => (__webpack_require__.O[key](chunkIds[j])))) {
/******/ 						chunkIds.splice(j--, 1);
/******/ 					} else {
/******/ 						fulfilled = false;
/******/ 						if(priority < notFulfilled) notFulfilled = priority;
/******/ 					}
/******/ 				}
/******/ 				if(fulfilled) {
/******/ 					deferred.splice(i--, 1)
/******/ 					var r = fn();
/******/ 					if (r !== undefined) result = r;
/******/ 				}
/******/ 			}
/******/ 			return result;
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/hasOwnProperty shorthand */
/******/ 	(() => {
/******/ 		__webpack_require__.o = (obj, prop) => (Object.prototype.hasOwnProperty.call(obj, prop))
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/make namespace object */
/******/ 	(() => {
/******/ 		// define __esModule on exports
/******/ 		__webpack_require__.r = (exports) => {
/******/ 			if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 				Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 			}
/******/ 			Object.defineProperty(exports, '__esModule', { value: true });
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/jsonp chunk loading */
/******/ 	(() => {
/******/ 		// no baseURI
/******/ 		
/******/ 		// object to store loaded and loading chunks
/******/ 		// undefined = chunk not loaded, null = chunk preloaded/prefetched
/******/ 		// [resolve, reject, Promise] = chunk loading, 0 = chunk loaded
/******/ 		var installedChunks = {
/******/ 			"/assets/js/app": 0,
/******/ 			"assets/css/style": 0
/******/ 		};
/******/ 		
/******/ 		// no chunk on demand loading
/******/ 		
/******/ 		// no prefetching
/******/ 		
/******/ 		// no preloaded
/******/ 		
/******/ 		// no HMR
/******/ 		
/******/ 		// no HMR manifest
/******/ 		
/******/ 		__webpack_require__.O.j = (chunkId) => (installedChunks[chunkId] === 0);
/******/ 		
/******/ 		// install a JSONP callback for chunk loading
/******/ 		var webpackJsonpCallback = (parentChunkLoadingFunction, data) => {
/******/ 			var [chunkIds, moreModules, runtime] = data;
/******/ 			// add "moreModules" to the modules object,
/******/ 			// then flag all "chunkIds" as loaded and fire callback
/******/ 			var moduleId, chunkId, i = 0;
/******/ 			if(chunkIds.some((id) => (installedChunks[id] !== 0))) {
/******/ 				for(moduleId in moreModules) {
/******/ 					if(__webpack_require__.o(moreModules, moduleId)) {
/******/ 						__webpack_require__.m[moduleId] = moreModules[moduleId];
/******/ 					}
/******/ 				}
/******/ 				if(runtime) var result = runtime(__webpack_require__);
/******/ 			}
/******/ 			if(parentChunkLoadingFunction) parentChunkLoadingFunction(data);
/******/ 			for(;i < chunkIds.length; i++) {
/******/ 				chunkId = chunkIds[i];
/******/ 				if(__webpack_require__.o(installedChunks, chunkId) && installedChunks[chunkId]) {
/******/ 					installedChunks[chunkId][0]();
/******/ 				}
/******/ 				installedChunks[chunkIds[i]] = 0;
/******/ 			}
/******/ 			return __webpack_require__.O(result);
/******/ 		}
/******/ 		
/******/ 		var chunkLoadingGlobal = self["webpackChunk"] = self["webpackChunk"] || [];
/******/ 		chunkLoadingGlobal.forEach(webpackJsonpCallback.bind(null, 0));
/******/ 		chunkLoadingGlobal.push = webpackJsonpCallback.bind(null, chunkLoadingGlobal.push.bind(chunkLoadingGlobal));
/******/ 	})();
/******/ 	
/************************************************************************/
/******/ 	
/******/ 	// startup
/******/ 	// Load entry module and return exports
/******/ 	// This entry module depends on other loaded chunks and execution need to be delayed
/******/ 	__webpack_require__.O(undefined, ["assets/css/style"], () => (__webpack_require__("./resources/js/detail_modal.js")))
/******/ 	__webpack_require__.O(undefined, ["assets/css/style"], () => (__webpack_require__("./resources/js/filter.js")))
/******/ 	__webpack_require__.O(undefined, ["assets/css/style"], () => (__webpack_require__("./resources/js/trim_names.js")))
/******/ 	var __webpack_exports__ = __webpack_require__.O(undefined, ["assets/css/style"], () => (__webpack_require__("./resources/sass/style.scss")))
/******/ 	__webpack_exports__ = __webpack_require__.O(__webpack_exports__);
/******/ 	
/******/ })()
;