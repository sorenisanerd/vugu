package vugu

// GENERATED FILE, DO NOT EDIT!  See renderer-js-script-maker.go

const jsHelperScript = "\n(function() {\n\n\tif (window.vuguRender) { return; } // only once\n\n    const opcodeEnd = 0         // no more instructions in this buffer\n    const opcodeClearRefmap = 1 // clear the reference map, all following instructions must not reference prior IDs\n    const opcodeSetHTMLRef = 2  // assign ref for html tag\n    const opcodeSetHeadRef = 3  // assign ref for head tag\n    const opcodeSetBodyRef = 4  // assign ref for body tag\n    const opcodeSelectRef = 5   // select element by ref\n    const opcodeSetAttrStr = 6  // assign attribute string to the current selected element\n\n    // Decoder provides our binary decoding.\n    // Using a class because that's what all the cool JS kids are doing these days.\n    class Decoder {\n\n        constructor(dataView, offset) {\n            this.dataView = dataView;\n            this.offset = offset || 0;\n            return this;\n        }\n\n        // readUint8 reads a single byte, 0-255\n        readUint8() {\n            var ret = this.dataView.getUint8(this.offset);\n            this.offset++;\n            return ret;\n        }\n\n        // readRefToString reads a 64-bit unsigned int ref but returns it as a hex string\n        readRefToString() {\n            // read in two 32-bit parts, BigInt is not yet well supported\n            var ret = this.dataView.getUint32(this.offset).toString(16).padStart(8, \"0\") +\n                this.dataView.getUint32(this.offset + 4).toString(16).padStart(8, \"0\");\n            this.offset += 8;\n            return ret;\n        }\n\n        // readString is 4 bytes length followed by utf chars\n        readString() {\n            var len = this.dataView.getUint32(this.offset);\n            var ret = utf8decoder.decode(new DataView(this.dataView.buffer, this.dataView.byteOffset + this.offset + 4, len));\n            this.offset += len + 4;\n            return ret;\n        }\n\n    }\n\n    let utf8decoder = new TextDecoder();\n\n    window._vuguRefMap = {};\n\n\twindow.vuguRender = function(buffer) {    \n\n\t\tlet state = window.vuguRenderState || {};\n\t\twindow.vuguRenderState = state;\n\n\t\tconsole.log(\"vuguRender called\", buffer);\n\n\t\tlet bufferView = new DataView(buffer.buffer, buffer.byteOffset, buffer.byteLength);\n\n        var decoder = new Decoder(bufferView, 0);\n        \n        var refmap = window._vuguRefMap;\n\n        var curref = \"\"; // current reference number\n        var currefel = null; // current reference element\n\n        instructionLoop: while (true) {\n\n\t\t\tlet opcode = decoder.readUint8();\n\n            switch (opcode) {\n\n                case opcodeEnd:\n                        break instructionLoop;\n    \n                case opcodeClearRefmap:\n                    refmap = {};\n                    window._vuguRefMap = refmap;\n                    curref = \"\";\n                    currefel = null;\n                    break;\n\n                case opcodeSetHTMLRef:\n                    var refstr = decoder.readRefToString();\n                    refmap[refstr] = document.querySelector(\"html\");\n                    break;\n\n                case opcodeSelectRef:\n                    var refstr = decoder.readRefToString();\n                    curref = refstr;\n                    currefel = refmap[refstr];\n                    if (!currefel) {\n                        console.error(\"opcodeSelectRef: refstr does not exist\", refstr);\n                    }\n                    break;\n\n                case opcodeSetAttrStr:\n                    if (!currefel) {\n                        console.error(\"opcodeSetAttrStr: no current reference\");\n                    }\n                    var attrName = decoder.readString();\n                    var attrValue = decoder.readString();\n                    currefel.setAttribute(attrName, attrValue);\n                    break;\n\n                default:\n                    console.error(\"found invalid opcode\", opcode);\n                    return;\n            }\n\n\t\t}\n\n\t}\n\n})()\n"