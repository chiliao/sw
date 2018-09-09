"""
/*
 * Copyright 2017-2018 Pensando Systems, Inc.  All rights reserved.
 *
 * This program is free software; you may redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; version 2 of the License.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
 * BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 * ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */
"""

DEV = { 
   "priority"                   : "default",
   "bus type"                   : "pci",
   "device info"                : { 
      "vendor list"                : [ 
         {   
            "vendor id"               : "1dd8",
            "vendor description"      : "Pensando Systems, Inc",
            "devices" : [ 
               {   
                  "device id"         : "1002",
                  "description"       : "Pensando Ethernet PF",
               }, 
               {   
                  "device id"         : "1003",
                  "description"       : "Pensando Ethernet VF",
               }, 
               {   
                  "device id"         : "1004",
                  "description"       : "Pensando Ethernet Management",
               }, 
            ],  
         },  
      ]   
   },  
}
