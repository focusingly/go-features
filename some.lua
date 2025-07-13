-- sample.lua
-- Lua 示例脚本，用于学习和调试

-- === 基础类型 ===
local number = 42
local str = "hello"
local flag = true
local nothing = nil

-- === 表（Table）操作 ===
local person = {
  name = "Alice",
  age = 30,
  skills = { "lua", "go", "js" }
}

function person:greet()
  print("Hi, my name is " .. self.name)
end

person.greet(person)

-- 遍历 skills 数组
for i, skill in ipairs(person.skills) do
  print("Skill[" .. i .. "] = " .. skill)
end

-- === 函数和返回值 ===
local function add(a, b)
  return a + b
end

local sum = add(10, 20)
print("10 + 20 = " .. sum)

-- 多返回值
function divide(a, b)
  if b == 0 then return nil, "division by zero" end
  return a / b
end

local res, err = divide(10, 0)
if not res then
  print("Error: " .. err)
end

-- === 表作为类/模块 ===
local Animal = {}
Animal.__index = Animal

function Animal:new(name)
  return setmetatable({name = name}, Animal)
end

function Animal:speak()
  print(self.name .. " makes a sound.")
end

local dog = Animal:new("Dog")
dog:speak()

-- === 协程（coroutine）===
local co = coroutine.create(function()
  for i = 1, 3 do
    print("Coroutine step " .. i)
    coroutine.yield()
  end
end)

coroutine.resume(co)
coroutine.resume(co)
coroutine.resume(co)

-- === 元表与运算符重载 ===
local Vector = {}
Vector.__index = Vector

function Vector:new(x, y)
  return setmetatable({x = x, y = y}, Vector)
end

function Vector:__add(v)
  return Vector:new(self.x + v.x, self.y + v.y)
end

function Vector:__tostring()
  return "(" .. self.x .. ", " .. self.y .. ")"
end

local v1 = Vector:new(1, 2)
local v2 = Vector:new(3, 4)
local v3 = v1 + v2
print("v1 + v2 = " .. v3)

-- === 错误处理 ===
local function risky()
  error("something went wrong")
end

local status, err = pcall(risky)
if not status then
  print("Caught error: " .. err)
end

-- === 返回值（用于 Go 调用时提取结果）===
return {
  number = number,
  person = person,
  sum = sum,
  vector = tostring(v3)
}
