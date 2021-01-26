/*
 Navicat Premium Data Transfer

 Source Server         : WebServer
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 24/01/2021 19:59:46
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for user_data
-- ----------------------------
DROP TABLE IF EXISTS "user_data";
CREATE TABLE "user_data" (
  "uid" text,
  "usernickname" TEXT,
  "userage" text,
  "usergender" TEXT,
  "usercontact" TEXT,
  "userbriefid" text
);

PRAGMA foreign_keys = true;
