/*eslint-disable block-scoped-var, no-redeclare, no-control-regex, no-prototype-builtins*/

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.irmaproto = (function() {

    /**
     * Namespace irmaproto.
     * @exports irmaproto
     * @namespace
     */
    var irmaproto = {};

    irmaproto.IRMASchemeMetadata = (function() {

        /**
         * Properties of a IRMASchemeMetadata.
         * @memberof irmaproto
         * @interface IIRMASchemeMetadata
         * @property {number|null} [version] IRMASchemeMetadata version
         * @property {string|null} [id] IRMASchemeMetadata id
         * @property {string|null} [url] IRMASchemeMetadata url
         * @property {Array.<irmaproto.IRMASchemeMetadata.ILocalizedName>|null} [name] IRMASchemeMetadata name
         * @property {Array.<irmaproto.IRMASchemeMetadata.ILocalizedDescription>|null} [description] IRMASchemeMetadata description
         * @property {string|null} [keyshareserver] IRMASchemeMetadata keyshareserver
         * @property {string|null} [keysharewebsite] IRMASchemeMetadata keysharewebsite
         * @property {string|null} [keyshareattribute] IRMASchemeMetadata keyshareattribute
         * @property {string|null} [contact] IRMASchemeMetadata contact
         */

        /**
         * Constructs a new IRMASchemeMetadata.
         * @memberof irmaproto
         * @classdesc Represents a IRMASchemeMetadata.
         * @implements IIRMASchemeMetadata
         * @constructor
         * @param {irmaproto.IIRMASchemeMetadata=} [properties] Properties to set
         */
        function IRMASchemeMetadata(properties) {
            this.name = [];
            this.description = [];
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * IRMASchemeMetadata version.
         * @member {number} version
         * @memberof irmaproto.IRMASchemeMetadata
         * @instance
         */
        IRMASchemeMetadata.prototype.version = 0;

        /**
         * IRMASchemeMetadata id.
         * @member {string} id
         * @memberof irmaproto.IRMASchemeMetadata
         * @instance
         */
        IRMASchemeMetadata.prototype.id = "";

        /**
         * IRMASchemeMetadata url.
         * @member {string} url
         * @memberof irmaproto.IRMASchemeMetadata
         * @instance
         */
        IRMASchemeMetadata.prototype.url = "";

        /**
         * IRMASchemeMetadata name.
         * @member {Array.<irmaproto.IRMASchemeMetadata.ILocalizedName>} name
         * @memberof irmaproto.IRMASchemeMetadata
         * @instance
         */
        IRMASchemeMetadata.prototype.name = $util.emptyArray;

        /**
         * IRMASchemeMetadata description.
         * @member {Array.<irmaproto.IRMASchemeMetadata.ILocalizedDescription>} description
         * @memberof irmaproto.IRMASchemeMetadata
         * @instance
         */
        IRMASchemeMetadata.prototype.description = $util.emptyArray;

        /**
         * IRMASchemeMetadata keyshareserver.
         * @member {string} keyshareserver
         * @memberof irmaproto.IRMASchemeMetadata
         * @instance
         */
        IRMASchemeMetadata.prototype.keyshareserver = "";

        /**
         * IRMASchemeMetadata keysharewebsite.
         * @member {string} keysharewebsite
         * @memberof irmaproto.IRMASchemeMetadata
         * @instance
         */
        IRMASchemeMetadata.prototype.keysharewebsite = "";

        /**
         * IRMASchemeMetadata keyshareattribute.
         * @member {string} keyshareattribute
         * @memberof irmaproto.IRMASchemeMetadata
         * @instance
         */
        IRMASchemeMetadata.prototype.keyshareattribute = "";

        /**
         * IRMASchemeMetadata contact.
         * @member {string} contact
         * @memberof irmaproto.IRMASchemeMetadata
         * @instance
         */
        IRMASchemeMetadata.prototype.contact = "";

        /**
         * Creates a new IRMASchemeMetadata instance using the specified properties.
         * @function create
         * @memberof irmaproto.IRMASchemeMetadata
         * @static
         * @param {irmaproto.IIRMASchemeMetadata=} [properties] Properties to set
         * @returns {irmaproto.IRMASchemeMetadata} IRMASchemeMetadata instance
         */
        IRMASchemeMetadata.create = function create(properties) {
            return new IRMASchemeMetadata(properties);
        };

        /**
         * Encodes the specified IRMASchemeMetadata message. Does not implicitly {@link irmaproto.IRMASchemeMetadata.verify|verify} messages.
         * @function encode
         * @memberof irmaproto.IRMASchemeMetadata
         * @static
         * @param {irmaproto.IIRMASchemeMetadata} message IRMASchemeMetadata message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        IRMASchemeMetadata.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.version != null && message.hasOwnProperty("version"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.version);
            if (message.id != null && message.hasOwnProperty("id"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.id);
            if (message.url != null && message.hasOwnProperty("url"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.url);
            if (message.name != null && message.name.length)
                for (var i = 0; i < message.name.length; ++i)
                    $root.irmaproto.IRMASchemeMetadata.LocalizedName.encode(message.name[i], writer.uint32(/* id 4, wireType 2 =*/34).fork()).ldelim();
            if (message.description != null && message.description.length)
                for (var i = 0; i < message.description.length; ++i)
                    $root.irmaproto.IRMASchemeMetadata.LocalizedDescription.encode(message.description[i], writer.uint32(/* id 5, wireType 2 =*/42).fork()).ldelim();
            if (message.keyshareserver != null && message.hasOwnProperty("keyshareserver"))
                writer.uint32(/* id 6, wireType 2 =*/50).string(message.keyshareserver);
            if (message.keysharewebsite != null && message.hasOwnProperty("keysharewebsite"))
                writer.uint32(/* id 7, wireType 2 =*/58).string(message.keysharewebsite);
            if (message.keyshareattribute != null && message.hasOwnProperty("keyshareattribute"))
                writer.uint32(/* id 8, wireType 2 =*/66).string(message.keyshareattribute);
            if (message.contact != null && message.hasOwnProperty("contact"))
                writer.uint32(/* id 9, wireType 2 =*/74).string(message.contact);
            return writer;
        };

        /**
         * Encodes the specified IRMASchemeMetadata message, length delimited. Does not implicitly {@link irmaproto.IRMASchemeMetadata.verify|verify} messages.
         * @function encodeDelimited
         * @memberof irmaproto.IRMASchemeMetadata
         * @static
         * @param {irmaproto.IIRMASchemeMetadata} message IRMASchemeMetadata message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        IRMASchemeMetadata.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a IRMASchemeMetadata message from the specified reader or buffer.
         * @function decode
         * @memberof irmaproto.IRMASchemeMetadata
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {irmaproto.IRMASchemeMetadata} IRMASchemeMetadata
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        IRMASchemeMetadata.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.irmaproto.IRMASchemeMetadata();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.version = reader.int32();
                    break;
                case 2:
                    message.id = reader.string();
                    break;
                case 3:
                    message.url = reader.string();
                    break;
                case 4:
                    if (!(message.name && message.name.length))
                        message.name = [];
                    message.name.push($root.irmaproto.IRMASchemeMetadata.LocalizedName.decode(reader, reader.uint32()));
                    break;
                case 5:
                    if (!(message.description && message.description.length))
                        message.description = [];
                    message.description.push($root.irmaproto.IRMASchemeMetadata.LocalizedDescription.decode(reader, reader.uint32()));
                    break;
                case 6:
                    message.keyshareserver = reader.string();
                    break;
                case 7:
                    message.keysharewebsite = reader.string();
                    break;
                case 8:
                    message.keyshareattribute = reader.string();
                    break;
                case 9:
                    message.contact = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a IRMASchemeMetadata message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof irmaproto.IRMASchemeMetadata
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {irmaproto.IRMASchemeMetadata} IRMASchemeMetadata
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        IRMASchemeMetadata.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a IRMASchemeMetadata message.
         * @function verify
         * @memberof irmaproto.IRMASchemeMetadata
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        IRMASchemeMetadata.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.version != null && message.hasOwnProperty("version"))
                if (!$util.isInteger(message.version))
                    return "version: integer expected";
            if (message.id != null && message.hasOwnProperty("id"))
                if (!$util.isString(message.id))
                    return "id: string expected";
            if (message.url != null && message.hasOwnProperty("url"))
                if (!$util.isString(message.url))
                    return "url: string expected";
            if (message.name != null && message.hasOwnProperty("name")) {
                if (!Array.isArray(message.name))
                    return "name: array expected";
                for (var i = 0; i < message.name.length; ++i) {
                    var error = $root.irmaproto.IRMASchemeMetadata.LocalizedName.verify(message.name[i]);
                    if (error)
                        return "name." + error;
                }
            }
            if (message.description != null && message.hasOwnProperty("description")) {
                if (!Array.isArray(message.description))
                    return "description: array expected";
                for (var i = 0; i < message.description.length; ++i) {
                    var error = $root.irmaproto.IRMASchemeMetadata.LocalizedDescription.verify(message.description[i]);
                    if (error)
                        return "description." + error;
                }
            }
            if (message.keyshareserver != null && message.hasOwnProperty("keyshareserver"))
                if (!$util.isString(message.keyshareserver))
                    return "keyshareserver: string expected";
            if (message.keysharewebsite != null && message.hasOwnProperty("keysharewebsite"))
                if (!$util.isString(message.keysharewebsite))
                    return "keysharewebsite: string expected";
            if (message.keyshareattribute != null && message.hasOwnProperty("keyshareattribute"))
                if (!$util.isString(message.keyshareattribute))
                    return "keyshareattribute: string expected";
            if (message.contact != null && message.hasOwnProperty("contact"))
                if (!$util.isString(message.contact))
                    return "contact: string expected";
            return null;
        };

        /**
         * Creates a IRMASchemeMetadata message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof irmaproto.IRMASchemeMetadata
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {irmaproto.IRMASchemeMetadata} IRMASchemeMetadata
         */
        IRMASchemeMetadata.fromObject = function fromObject(object) {
            if (object instanceof $root.irmaproto.IRMASchemeMetadata)
                return object;
            var message = new $root.irmaproto.IRMASchemeMetadata();
            if (object.version != null)
                message.version = object.version | 0;
            if (object.id != null)
                message.id = String(object.id);
            if (object.url != null)
                message.url = String(object.url);
            if (object.name) {
                if (!Array.isArray(object.name))
                    throw TypeError(".irmaproto.IRMASchemeMetadata.name: array expected");
                message.name = [];
                for (var i = 0; i < object.name.length; ++i) {
                    if (typeof object.name[i] !== "object")
                        throw TypeError(".irmaproto.IRMASchemeMetadata.name: object expected");
                    message.name[i] = $root.irmaproto.IRMASchemeMetadata.LocalizedName.fromObject(object.name[i]);
                }
            }
            if (object.description) {
                if (!Array.isArray(object.description))
                    throw TypeError(".irmaproto.IRMASchemeMetadata.description: array expected");
                message.description = [];
                for (var i = 0; i < object.description.length; ++i) {
                    if (typeof object.description[i] !== "object")
                        throw TypeError(".irmaproto.IRMASchemeMetadata.description: object expected");
                    message.description[i] = $root.irmaproto.IRMASchemeMetadata.LocalizedDescription.fromObject(object.description[i]);
                }
            }
            if (object.keyshareserver != null)
                message.keyshareserver = String(object.keyshareserver);
            if (object.keysharewebsite != null)
                message.keysharewebsite = String(object.keysharewebsite);
            if (object.keyshareattribute != null)
                message.keyshareattribute = String(object.keyshareattribute);
            if (object.contact != null)
                message.contact = String(object.contact);
            return message;
        };

        /**
         * Creates a plain object from a IRMASchemeMetadata message. Also converts values to other types if specified.
         * @function toObject
         * @memberof irmaproto.IRMASchemeMetadata
         * @static
         * @param {irmaproto.IRMASchemeMetadata} message IRMASchemeMetadata
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        IRMASchemeMetadata.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.arrays || options.defaults) {
                object.name = [];
                object.description = [];
            }
            if (options.defaults) {
                object.version = 0;
                object.id = "";
                object.url = "";
                object.keyshareserver = "";
                object.keysharewebsite = "";
                object.keyshareattribute = "";
                object.contact = "";
            }
            if (message.version != null && message.hasOwnProperty("version"))
                object.version = message.version;
            if (message.id != null && message.hasOwnProperty("id"))
                object.id = message.id;
            if (message.url != null && message.hasOwnProperty("url"))
                object.url = message.url;
            if (message.name && message.name.length) {
                object.name = [];
                for (var j = 0; j < message.name.length; ++j)
                    object.name[j] = $root.irmaproto.IRMASchemeMetadata.LocalizedName.toObject(message.name[j], options);
            }
            if (message.description && message.description.length) {
                object.description = [];
                for (var j = 0; j < message.description.length; ++j)
                    object.description[j] = $root.irmaproto.IRMASchemeMetadata.LocalizedDescription.toObject(message.description[j], options);
            }
            if (message.keyshareserver != null && message.hasOwnProperty("keyshareserver"))
                object.keyshareserver = message.keyshareserver;
            if (message.keysharewebsite != null && message.hasOwnProperty("keysharewebsite"))
                object.keysharewebsite = message.keysharewebsite;
            if (message.keyshareattribute != null && message.hasOwnProperty("keyshareattribute"))
                object.keyshareattribute = message.keyshareattribute;
            if (message.contact != null && message.hasOwnProperty("contact"))
                object.contact = message.contact;
            return object;
        };

        /**
         * Converts this IRMASchemeMetadata to JSON.
         * @function toJSON
         * @memberof irmaproto.IRMASchemeMetadata
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        IRMASchemeMetadata.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        IRMASchemeMetadata.LocalizedName = (function() {

            /**
             * Properties of a LocalizedName.
             * @memberof irmaproto.IRMASchemeMetadata
             * @interface ILocalizedName
             * @property {string|null} [lang] LocalizedName lang
             * @property {string|null} [name] LocalizedName name
             */

            /**
             * Constructs a new LocalizedName.
             * @memberof irmaproto.IRMASchemeMetadata
             * @classdesc Represents a LocalizedName.
             * @implements ILocalizedName
             * @constructor
             * @param {irmaproto.IRMASchemeMetadata.ILocalizedName=} [properties] Properties to set
             */
            function LocalizedName(properties) {
                if (properties)
                    for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * LocalizedName lang.
             * @member {string} lang
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedName
             * @instance
             */
            LocalizedName.prototype.lang = "";

            /**
             * LocalizedName name.
             * @member {string} name
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedName
             * @instance
             */
            LocalizedName.prototype.name = "";

            /**
             * Creates a new LocalizedName instance using the specified properties.
             * @function create
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedName
             * @static
             * @param {irmaproto.IRMASchemeMetadata.ILocalizedName=} [properties] Properties to set
             * @returns {irmaproto.IRMASchemeMetadata.LocalizedName} LocalizedName instance
             */
            LocalizedName.create = function create(properties) {
                return new LocalizedName(properties);
            };

            /**
             * Encodes the specified LocalizedName message. Does not implicitly {@link irmaproto.IRMASchemeMetadata.LocalizedName.verify|verify} messages.
             * @function encode
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedName
             * @static
             * @param {irmaproto.IRMASchemeMetadata.ILocalizedName} message LocalizedName message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            LocalizedName.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.lang != null && message.hasOwnProperty("lang"))
                    writer.uint32(/* id 1, wireType 2 =*/10).string(message.lang);
                if (message.name != null && message.hasOwnProperty("name"))
                    writer.uint32(/* id 2, wireType 2 =*/18).string(message.name);
                return writer;
            };

            /**
             * Encodes the specified LocalizedName message, length delimited. Does not implicitly {@link irmaproto.IRMASchemeMetadata.LocalizedName.verify|verify} messages.
             * @function encodeDelimited
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedName
             * @static
             * @param {irmaproto.IRMASchemeMetadata.ILocalizedName} message LocalizedName message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            LocalizedName.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a LocalizedName message from the specified reader or buffer.
             * @function decode
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedName
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {irmaproto.IRMASchemeMetadata.LocalizedName} LocalizedName
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            LocalizedName.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                var end = length === undefined ? reader.len : reader.pos + length, message = new $root.irmaproto.IRMASchemeMetadata.LocalizedName();
                while (reader.pos < end) {
                    var tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.lang = reader.string();
                        break;
                    case 2:
                        message.name = reader.string();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a LocalizedName message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedName
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {irmaproto.IRMASchemeMetadata.LocalizedName} LocalizedName
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            LocalizedName.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a LocalizedName message.
             * @function verify
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedName
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            LocalizedName.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.lang != null && message.hasOwnProperty("lang"))
                    if (!$util.isString(message.lang))
                        return "lang: string expected";
                if (message.name != null && message.hasOwnProperty("name"))
                    if (!$util.isString(message.name))
                        return "name: string expected";
                return null;
            };

            /**
             * Creates a LocalizedName message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedName
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {irmaproto.IRMASchemeMetadata.LocalizedName} LocalizedName
             */
            LocalizedName.fromObject = function fromObject(object) {
                if (object instanceof $root.irmaproto.IRMASchemeMetadata.LocalizedName)
                    return object;
                var message = new $root.irmaproto.IRMASchemeMetadata.LocalizedName();
                if (object.lang != null)
                    message.lang = String(object.lang);
                if (object.name != null)
                    message.name = String(object.name);
                return message;
            };

            /**
             * Creates a plain object from a LocalizedName message. Also converts values to other types if specified.
             * @function toObject
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedName
             * @static
             * @param {irmaproto.IRMASchemeMetadata.LocalizedName} message LocalizedName
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            LocalizedName.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                var object = {};
                if (options.defaults) {
                    object.lang = "";
                    object.name = "";
                }
                if (message.lang != null && message.hasOwnProperty("lang"))
                    object.lang = message.lang;
                if (message.name != null && message.hasOwnProperty("name"))
                    object.name = message.name;
                return object;
            };

            /**
             * Converts this LocalizedName to JSON.
             * @function toJSON
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedName
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            LocalizedName.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return LocalizedName;
        })();

        IRMASchemeMetadata.LocalizedDescription = (function() {

            /**
             * Properties of a LocalizedDescription.
             * @memberof irmaproto.IRMASchemeMetadata
             * @interface ILocalizedDescription
             * @property {string|null} [lang] LocalizedDescription lang
             * @property {string|null} [name] LocalizedDescription name
             */

            /**
             * Constructs a new LocalizedDescription.
             * @memberof irmaproto.IRMASchemeMetadata
             * @classdesc Represents a LocalizedDescription.
             * @implements ILocalizedDescription
             * @constructor
             * @param {irmaproto.IRMASchemeMetadata.ILocalizedDescription=} [properties] Properties to set
             */
            function LocalizedDescription(properties) {
                if (properties)
                    for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * LocalizedDescription lang.
             * @member {string} lang
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedDescription
             * @instance
             */
            LocalizedDescription.prototype.lang = "";

            /**
             * LocalizedDescription name.
             * @member {string} name
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedDescription
             * @instance
             */
            LocalizedDescription.prototype.name = "";

            /**
             * Creates a new LocalizedDescription instance using the specified properties.
             * @function create
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedDescription
             * @static
             * @param {irmaproto.IRMASchemeMetadata.ILocalizedDescription=} [properties] Properties to set
             * @returns {irmaproto.IRMASchemeMetadata.LocalizedDescription} LocalizedDescription instance
             */
            LocalizedDescription.create = function create(properties) {
                return new LocalizedDescription(properties);
            };

            /**
             * Encodes the specified LocalizedDescription message. Does not implicitly {@link irmaproto.IRMASchemeMetadata.LocalizedDescription.verify|verify} messages.
             * @function encode
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedDescription
             * @static
             * @param {irmaproto.IRMASchemeMetadata.ILocalizedDescription} message LocalizedDescription message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            LocalizedDescription.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.lang != null && message.hasOwnProperty("lang"))
                    writer.uint32(/* id 1, wireType 2 =*/10).string(message.lang);
                if (message.name != null && message.hasOwnProperty("name"))
                    writer.uint32(/* id 2, wireType 2 =*/18).string(message.name);
                return writer;
            };

            /**
             * Encodes the specified LocalizedDescription message, length delimited. Does not implicitly {@link irmaproto.IRMASchemeMetadata.LocalizedDescription.verify|verify} messages.
             * @function encodeDelimited
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedDescription
             * @static
             * @param {irmaproto.IRMASchemeMetadata.ILocalizedDescription} message LocalizedDescription message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            LocalizedDescription.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a LocalizedDescription message from the specified reader or buffer.
             * @function decode
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedDescription
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {irmaproto.IRMASchemeMetadata.LocalizedDescription} LocalizedDescription
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            LocalizedDescription.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                var end = length === undefined ? reader.len : reader.pos + length, message = new $root.irmaproto.IRMASchemeMetadata.LocalizedDescription();
                while (reader.pos < end) {
                    var tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.lang = reader.string();
                        break;
                    case 2:
                        message.name = reader.string();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a LocalizedDescription message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedDescription
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {irmaproto.IRMASchemeMetadata.LocalizedDescription} LocalizedDescription
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            LocalizedDescription.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a LocalizedDescription message.
             * @function verify
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedDescription
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            LocalizedDescription.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.lang != null && message.hasOwnProperty("lang"))
                    if (!$util.isString(message.lang))
                        return "lang: string expected";
                if (message.name != null && message.hasOwnProperty("name"))
                    if (!$util.isString(message.name))
                        return "name: string expected";
                return null;
            };

            /**
             * Creates a LocalizedDescription message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedDescription
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {irmaproto.IRMASchemeMetadata.LocalizedDescription} LocalizedDescription
             */
            LocalizedDescription.fromObject = function fromObject(object) {
                if (object instanceof $root.irmaproto.IRMASchemeMetadata.LocalizedDescription)
                    return object;
                var message = new $root.irmaproto.IRMASchemeMetadata.LocalizedDescription();
                if (object.lang != null)
                    message.lang = String(object.lang);
                if (object.name != null)
                    message.name = String(object.name);
                return message;
            };

            /**
             * Creates a plain object from a LocalizedDescription message. Also converts values to other types if specified.
             * @function toObject
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedDescription
             * @static
             * @param {irmaproto.IRMASchemeMetadata.LocalizedDescription} message LocalizedDescription
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            LocalizedDescription.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                var object = {};
                if (options.defaults) {
                    object.lang = "";
                    object.name = "";
                }
                if (message.lang != null && message.hasOwnProperty("lang"))
                    object.lang = message.lang;
                if (message.name != null && message.hasOwnProperty("name"))
                    object.name = message.name;
                return object;
            };

            /**
             * Converts this LocalizedDescription to JSON.
             * @function toJSON
             * @memberof irmaproto.IRMASchemeMetadata.LocalizedDescription
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            LocalizedDescription.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return LocalizedDescription;
        })();

        return IRMASchemeMetadata;
    })();

    return irmaproto;
})();

module.exports = $root;
