#! /usr/bin/env python

from zplot import *

# populate zplot table from data file
t = table('line.data')

canvas = pdf('line.pdf', dimensions=['3in','2in'])

# a drawable is a region of a canvas, and is used to convert data
# coordinates to raw pixel coordinates on the canvas, based on
# xrange and yrange.
d = drawable(canvas, xrange=[0,1024], yrange=[0,10])

# create an axis for our drawable.  Specify axis labels with
# [START,END,STEP] via xauto and yauto.
axis(d, xtitle='File Size (mb)', xauto=[0,1024,256],
     ytitle='Copy Time (s)', yauto=[0,10,2], ytitleshift=[0,0])

# a plotter fetches coordinates from the table, querying for the
# specified xfield and yfield, and converting these via the drawable.
p = plotter()

# each thing we plot can be added to the legend
L = legend()

# for line color, we use a string that is 'red,green,blue'.
# A primary color can range from 0 to 1, so '1,1,1' is white,
# '0,0,0' is black, and '1,0,0' is red.
p.line(d, t, xfield='size', yfield='disk_time',
       linewidth=2, linecolor='0.6,0.6,0.6',
       legend=L, legendtext='disk')

# for the linedash pattern, we specify [LINE,SPACE].  So in this
# case, the line will consist of 3-pixel dashes, each separated by
# a 1-pixel space.
p.line(d, t, xfield='size', yfield='ssd_time',
       linedash=[3,1], legend=L, legendtext='SSD')

# draw legend near top-left of drawable
L.draw(canvas, coord=[d.left()+20, d.top()-10])

# save to eps file
canvas.render()
