/**
 * Hippo.Web
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.HippoWeb);
  }
}(this, function(expect, HippoWeb) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new HippoWeb.ChannelJobStatusItemPage();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('ChannelJobStatusItemPage', function() {
    it('should create an instance of ChannelJobStatusItemPage', function() {
      // uncomment below and update the code to test ChannelJobStatusItemPage
      //var instance = new HippoWeb.ChannelJobStatusItemPage();
      //expect(instance).to.be.a(HippoWeb.ChannelJobStatusItemPage);
    });

    it('should have the property items (base name: "items")', function() {
      // uncomment below and update the code to test the property items
      //var instance = new HippoWeb.ChannelJobStatusItemPage();
      //expect(instance).to.be();
    });

    it('should have the property totalItems (base name: "totalItems")', function() {
      // uncomment below and update the code to test the property totalItems
      //var instance = new HippoWeb.ChannelJobStatusItemPage();
      //expect(instance).to.be();
    });

    it('should have the property pageIndex (base name: "pageIndex")', function() {
      // uncomment below and update the code to test the property pageIndex
      //var instance = new HippoWeb.ChannelJobStatusItemPage();
      //expect(instance).to.be();
    });

    it('should have the property pageSize (base name: "pageSize")', function() {
      // uncomment below and update the code to test the property pageSize
      //var instance = new HippoWeb.ChannelJobStatusItemPage();
      //expect(instance).to.be();
    });

    it('should have the property isLastPage (base name: "isLastPage")', function() {
      // uncomment below and update the code to test the property isLastPage
      //var instance = new HippoWeb.ChannelJobStatusItemPage();
      //expect(instance).to.be();
    });

  });

}));
