
-- SET THESE TO VARIABLES TO YOUR BBS -----------------------------------------
local bbsdomain = "mycoolbbs.com"  -- change to your domain name
local bbsdir = "/home/robbiew/bbs"   -- Windows requires double slash!
-------------------------------------------------------------------------------

local node = math.floor(bbs_get_node()+0.5) -- convert float to integer
local bbsname = bbs_get_bbs_name()
local sysopname = bbs_get_sysop_name()
local username = bbs_get_username()
local seclevel =  bbs_get_user_attribute("seclevel", "0")
local timeleft =  bbs_get_user_attribute("time_left", "0")
local cols =  math.floor(bbs_get_term_width()+0.5)  -- convert float to integer
local rows =  math.floor(bbs_get_term_height()+0.5) -- convert float to integer
local ostype = bbs_get_os()
local termtype =  string.lower(bbs_get_term_type())
local terminal = ""
local loadablefonts = false
local xtendpalette = false

-- detect term capabilities
if (termtype == string.lower("ANSI-256COLOR-RGB")) or (cols>80) then
    terminal = "Netrunner"
elseif termtype == string.lower("syncterm") then
    terminal = "Syncterm"
elseif termtype == string.lower("magiterm") then
    terminal = "Magiterm"
else
    terminal = "ANSI-Term"
end
if (terminal == "Netrunner") or (terminal == "ANSI-Term") or (terminal == "Magiterm")then
    loadablefonts = false
else
    loadablefonts = true
end
if (terminal == "Syncterm") or (terminal == "Netrunner") or (terminal == "Magiterm") then
    xtendpalette = true
else
    xtendpalette = false
end

-- write the dropfile

local path
if string.find(tostring(ostype), "windows") then
    path = (bbsdir .. "\\temp\\" .. node .. "\\phenomdrop.txt")
else
    path = (bbsdir .. "/temp/" .. node .. "/phenomdrop.txt")
end

local f, err = io.open(path, "w+")
bbs_write_string(err)

if f then
    f:write(node .. "\n")
    f:write(bbsname .. "\n")
    f:write(sysopname .. "\n")
    f:write(username .. "\n")
    f:write(seclevel .. "\n")
    f:write(timeleft .. "\n")
    f:write(cols .. "\n")
    f:write(rows .. "\n")
    f:write(ostype .. "\n")
    f:write(bbsdir .. "\n")
    f:write(bbsdomain .. "\n")
    f:write(tostring(loadablefonts) .. "\n")
    f:write(tostring(xtendpalette) .. "\n")
    f:close()
else
    bbs_write_string("Something is wrong with the drop file!")
end

-- done!