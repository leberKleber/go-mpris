# go-mpris
go mpris version 2.2 implementation 

##Functions:

###Player
https://specifications.freedesktop.org/mpris-spec/2.2/Player_Interface.html
####Methods:
- [X] Next
- [ ] Previous
- [ ] Pause
- [ ] PlayPause
- [ ] Stop
- [ ] Play
- [ ] Seek
- [ ] SetPosition
- [ ] OpenUri




---
Signals
Seeked 	(x: Position) 	

Properties
PlaybackStatus 	s (Playback_Status) 	Read only 		
LoopStatus 	s (Loop_Status) 	Read/Write 		(optional)
Rate 	d (Playback_Rate) 	Read/Write 		
Shuffle 	b 	Read/Write 		(optional)
Metadata 	a{sv} (Metadata_Map) 	Read only 		
Volume 	d (Volume) 	Read/Write 		
Position 	x (Time_In_Us) 	Read only 		
MinimumRate 	d (Playback_Rate) 	Read only 		
MaximumRate 	d (Playback_Rate) 	Read only 		
CanGoNext 	b 	Read only 		
CanGoPrevious 	b 	Read only 		
CanPlay 	b 	Read only 		
CanPause 	b 	Read only 		
CanSeek 	b 	Read only 		
CanControl 	b 	Read only 	

Types
Track_Id 	Simple Type 	o 	
Playback_Rate 	Simple Type 	d 	
Volume 	Simple Type 	d 	
Time_In_Us 	Simple Type 	x 	
Playback_Status 	Enum 	s 	
Loop_Status 	Enum 	s 	