/*
 Navicat Premium Data Transfer

 Source Server         : WebServer
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 24/01/2021 19:59:35
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "users";
CREATE TABLE "users" (
  "uid" text,
  "user_name" TEXT,
  "user_password" TEXT
);

PRAGMA foreign_keys = true;
