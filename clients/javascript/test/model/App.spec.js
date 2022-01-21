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
    instance = new HippoWeb.App();
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

  describe('App', function() {
    it('should create an instance of App', function() {
      // uncomment below and update the code to test App
      //var instance = new HippoWeb.App();
      //expect(instance).to.be.a(HippoWeb.App);
    });

    it('should have the property created (base name: "created")', function() {
      // uncomment below and update the code to test the property created
      //var instance = new HippoWeb.App();
      //expect(instance).to.be();
    });

    it('should have the property createdBy (base name: "createdBy")', function() {
      // uncomment below and update the code to test the property createdBy
      //var instance = new HippoWeb.App();
      //expect(instance).to.be();
    });

    it('should have the property lastModified (base name: "lastModified")', function() {
      // uncomment below and update the code to test the property lastModified
      //var instance = new HippoWeb.App();
      //expect(instance).to.be();
    });

    it('should have the property lastModifiedBy (base name: "lastModifiedBy")', function() {
      // uncomment below and update the code to test the property lastModifiedBy
      //var instance = new HippoWeb.App();
      //expect(instance).to.be();
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new HippoWeb.App();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new HippoWeb.App();
      //expect(instance).to.be();
    });

    it('should have the property storageId (base name: "storageId")', function() {
      // uncomment below and update the code to test the property storageId
      //var instance = new HippoWeb.App();
      //expect(instance).to.be();
    });

    it('should have the property channels (base name: "channels")', function() {
      // uncomment below and update the code to test the property channels
      //var instance = new HippoWeb.App();
      //expect(instance).to.be();
    });

    it('should have the property revisions (base name: "revisions")', function() {
      // uncomment below and update the code to test the property revisions
      //var instance = new HippoWeb.App();
      //expect(instance).to.be();
    });

    it('should have the property domainEvents (base name: "domainEvents")', function() {
      // uncomment below and update the code to test the property domainEvents
      //var instance = new HippoWeb.App();
      //expect(instance).to.be();
    });

  });

}));