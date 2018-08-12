// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_referee.proto

package sslproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// These are the "coarse" stages of the game.
type SSL_Referee_Stage int32

const (
	// The first half is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	SSL_Referee_NORMAL_FIRST_HALF_PRE SSL_Referee_Stage = 0
	// The first half of the normal game, before half time.
	SSL_Referee_NORMAL_FIRST_HALF SSL_Referee_Stage = 1
	// Half time between first and second halves.
	SSL_Referee_NORMAL_HALF_TIME SSL_Referee_Stage = 2
	// The second half is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	SSL_Referee_NORMAL_SECOND_HALF_PRE SSL_Referee_Stage = 3
	// The second half of the normal game, after half time.
	SSL_Referee_NORMAL_SECOND_HALF SSL_Referee_Stage = 4
	// The break before extra time.
	SSL_Referee_EXTRA_TIME_BREAK SSL_Referee_Stage = 5
	// The first half of extra time is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	SSL_Referee_EXTRA_FIRST_HALF_PRE SSL_Referee_Stage = 6
	// The first half of extra time.
	SSL_Referee_EXTRA_FIRST_HALF SSL_Referee_Stage = 7
	// Half time between first and second extra halves.
	SSL_Referee_EXTRA_HALF_TIME SSL_Referee_Stage = 8
	// The second half of extra time is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	SSL_Referee_EXTRA_SECOND_HALF_PRE SSL_Referee_Stage = 9
	// The second half of extra time.
	SSL_Referee_EXTRA_SECOND_HALF SSL_Referee_Stage = 10
	// The break before penalty shootout.
	SSL_Referee_PENALTY_SHOOTOUT_BREAK SSL_Referee_Stage = 11
	// The penalty shootout.
	SSL_Referee_PENALTY_SHOOTOUT SSL_Referee_Stage = 12
	// The game is over.
	SSL_Referee_POST_GAME SSL_Referee_Stage = 13
)

var SSL_Referee_Stage_name = map[int32]string{
	0:  "NORMAL_FIRST_HALF_PRE",
	1:  "NORMAL_FIRST_HALF",
	2:  "NORMAL_HALF_TIME",
	3:  "NORMAL_SECOND_HALF_PRE",
	4:  "NORMAL_SECOND_HALF",
	5:  "EXTRA_TIME_BREAK",
	6:  "EXTRA_FIRST_HALF_PRE",
	7:  "EXTRA_FIRST_HALF",
	8:  "EXTRA_HALF_TIME",
	9:  "EXTRA_SECOND_HALF_PRE",
	10: "EXTRA_SECOND_HALF",
	11: "PENALTY_SHOOTOUT_BREAK",
	12: "PENALTY_SHOOTOUT",
	13: "POST_GAME",
}
var SSL_Referee_Stage_value = map[string]int32{
	"NORMAL_FIRST_HALF_PRE":  0,
	"NORMAL_FIRST_HALF":      1,
	"NORMAL_HALF_TIME":       2,
	"NORMAL_SECOND_HALF_PRE": 3,
	"NORMAL_SECOND_HALF":     4,
	"EXTRA_TIME_BREAK":       5,
	"EXTRA_FIRST_HALF_PRE":   6,
	"EXTRA_FIRST_HALF":       7,
	"EXTRA_HALF_TIME":        8,
	"EXTRA_SECOND_HALF_PRE":  9,
	"EXTRA_SECOND_HALF":      10,
	"PENALTY_SHOOTOUT_BREAK": 11,
	"PENALTY_SHOOTOUT":       12,
	"POST_GAME":              13,
}

func (x SSL_Referee_Stage) Enum() *SSL_Referee_Stage {
	p := new(SSL_Referee_Stage)
	*p = x
	return p
}
func (x SSL_Referee_Stage) String() string {
	return proto.EnumName(SSL_Referee_Stage_name, int32(x))
}
func (x *SSL_Referee_Stage) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SSL_Referee_Stage_value, data, "SSL_Referee_Stage")
	if err != nil {
		return err
	}
	*x = SSL_Referee_Stage(value)
	return nil
}
func (SSL_Referee_Stage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ssl_referee_7314bdac51db39bf, []int{0, 0}
}

// These are the "fine" states of play on the field.
type SSL_Referee_Command int32

const (
	// All robots should completely stop moving.
	SSL_Referee_HALT SSL_Referee_Command = 0
	// Robots must keep 50 cm from the ball.
	SSL_Referee_STOP SSL_Referee_Command = 1
	// A prepared kickoff or penalty may now be taken.
	SSL_Referee_NORMAL_START SSL_Referee_Command = 2
	// The ball is dropped and free for either team.
	SSL_Referee_FORCE_START SSL_Referee_Command = 3
	// The yellow team may move into kickoff position.
	SSL_Referee_PREPARE_KICKOFF_YELLOW SSL_Referee_Command = 4
	// The blue team may move into kickoff position.
	SSL_Referee_PREPARE_KICKOFF_BLUE SSL_Referee_Command = 5
	// The yellow team may move into penalty position.
	SSL_Referee_PREPARE_PENALTY_YELLOW SSL_Referee_Command = 6
	// The blue team may move into penalty position.
	SSL_Referee_PREPARE_PENALTY_BLUE SSL_Referee_Command = 7
	// The yellow team may take a direct free kick.
	SSL_Referee_DIRECT_FREE_YELLOW SSL_Referee_Command = 8
	// The blue team may take a direct free kick.
	SSL_Referee_DIRECT_FREE_BLUE SSL_Referee_Command = 9
	// The yellow team may take an indirect free kick.
	SSL_Referee_INDIRECT_FREE_YELLOW SSL_Referee_Command = 10
	// The blue team may take an indirect free kick.
	SSL_Referee_INDIRECT_FREE_BLUE SSL_Referee_Command = 11
	// The yellow team is currently in a timeout.
	SSL_Referee_TIMEOUT_YELLOW SSL_Referee_Command = 12
	// The blue team is currently in a timeout.
	SSL_Referee_TIMEOUT_BLUE SSL_Referee_Command = 13
	// The yellow team just scored a goal.
	// For information only.
	// For rules compliance, teams must treat as STOP.
	SSL_Referee_GOAL_YELLOW SSL_Referee_Command = 14
	// The blue team just scored a goal.
	SSL_Referee_GOAL_BLUE SSL_Referee_Command = 15
	// Equivalent to STOP, but the yellow team must pick up the ball and
	// drop it in the Designated Position.
	SSL_Referee_BALL_PLACEMENT_YELLOW SSL_Referee_Command = 16
	// Equivalent to STOP, but the blue team must pick up the ball and drop
	// it in the Designated Position.
	SSL_Referee_BALL_PLACEMENT_BLUE SSL_Referee_Command = 17
)

var SSL_Referee_Command_name = map[int32]string{
	0:  "HALT",
	1:  "STOP",
	2:  "NORMAL_START",
	3:  "FORCE_START",
	4:  "PREPARE_KICKOFF_YELLOW",
	5:  "PREPARE_KICKOFF_BLUE",
	6:  "PREPARE_PENALTY_YELLOW",
	7:  "PREPARE_PENALTY_BLUE",
	8:  "DIRECT_FREE_YELLOW",
	9:  "DIRECT_FREE_BLUE",
	10: "INDIRECT_FREE_YELLOW",
	11: "INDIRECT_FREE_BLUE",
	12: "TIMEOUT_YELLOW",
	13: "TIMEOUT_BLUE",
	14: "GOAL_YELLOW",
	15: "GOAL_BLUE",
	16: "BALL_PLACEMENT_YELLOW",
	17: "BALL_PLACEMENT_BLUE",
}
var SSL_Referee_Command_value = map[string]int32{
	"HALT":                   0,
	"STOP":                   1,
	"NORMAL_START":           2,
	"FORCE_START":            3,
	"PREPARE_KICKOFF_YELLOW": 4,
	"PREPARE_KICKOFF_BLUE":   5,
	"PREPARE_PENALTY_YELLOW": 6,
	"PREPARE_PENALTY_BLUE":   7,
	"DIRECT_FREE_YELLOW":     8,
	"DIRECT_FREE_BLUE":       9,
	"INDIRECT_FREE_YELLOW":   10,
	"INDIRECT_FREE_BLUE":     11,
	"TIMEOUT_YELLOW":         12,
	"TIMEOUT_BLUE":           13,
	"GOAL_YELLOW":            14,
	"GOAL_BLUE":              15,
	"BALL_PLACEMENT_YELLOW":  16,
	"BALL_PLACEMENT_BLUE":    17,
}

func (x SSL_Referee_Command) Enum() *SSL_Referee_Command {
	p := new(SSL_Referee_Command)
	*p = x
	return p
}
func (x SSL_Referee_Command) String() string {
	return proto.EnumName(SSL_Referee_Command_name, int32(x))
}
func (x *SSL_Referee_Command) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SSL_Referee_Command_value, data, "SSL_Referee_Command")
	if err != nil {
		return err
	}
	*x = SSL_Referee_Command(value)
	return nil
}
func (SSL_Referee_Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ssl_referee_7314bdac51db39bf, []int{0, 1}
}

// Each UDP packet contains one of these messages.
type SSL_Referee struct {
	// The UNIX timestamp when the packet was sent, in microseconds.
	// Divide by 1,000,000 to get a time_t.
	PacketTimestamp *uint64            `protobuf:"varint,1,req,name=packet_timestamp,json=packetTimestamp" json:"packet_timestamp,omitempty"`
	Stage           *SSL_Referee_Stage `protobuf:"varint,2,req,name=stage,enum=SSL_Referee_Stage" json:"stage,omitempty"`
	// The number of microseconds left in the stage.
	// The following stages have this value; the rest do not:
	// NORMAL_FIRST_HALF
	// NORMAL_HALF_TIME
	// NORMAL_SECOND_HALF
	// EXTRA_TIME_BREAK
	// EXTRA_FIRST_HALF
	// EXTRA_HALF_TIME
	// EXTRA_SECOND_HALF
	// PENALTY_SHOOTOUT_BREAK
	//
	// If the stage runs over its specified time, this value
	// becomes negative.
	StageTimeLeft *int32               `protobuf:"zigzag32,3,opt,name=stage_time_left,json=stageTimeLeft" json:"stage_time_left,omitempty"`
	Command       *SSL_Referee_Command `protobuf:"varint,4,req,name=command,enum=SSL_Referee_Command" json:"command,omitempty"`
	// The number of commands issued since startup (mod 2^32).
	CommandCounter *uint32 `protobuf:"varint,5,req,name=command_counter,json=commandCounter" json:"command_counter,omitempty"`
	// The UNIX timestamp when the command was issued, in microseconds.
	// This value changes only when a new command is issued, not on each packet.
	CommandTimestamp *uint64 `protobuf:"varint,6,req,name=command_timestamp,json=commandTimestamp" json:"command_timestamp,omitempty"`
	// Information about the two teams.
	Yellow             *SSL_Referee_TeamInfo `protobuf:"bytes,7,req,name=yellow" json:"yellow,omitempty"`
	Blue               *SSL_Referee_TeamInfo `protobuf:"bytes,8,req,name=blue" json:"blue,omitempty"`
	DesignatedPosition *SSL_Referee_Point    `protobuf:"bytes,9,opt,name=designated_position,json=designatedPosition" json:"designated_position,omitempty"`
	// Information about the direction of play.
	// True, if the blue team will have it's goal on the positive x-axis of the ssl-vision coordinate system
	// Obviously, the yellow team will play on the opposide half
	// For compatibility, this field is optional
	BlueTeamOnPositiveHalf *bool `protobuf:"varint,10,opt,name=blueTeamOnPositiveHalf" json:"blueTeamOnPositiveHalf,omitempty"`
	// The game event that caused the referee command
	GameEvent            *SSL_Referee_Game_Event `protobuf:"bytes,11,opt,name=gameEvent" json:"gameEvent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SSL_Referee) Reset()         { *m = SSL_Referee{} }
func (m *SSL_Referee) String() string { return proto.CompactTextString(m) }
func (*SSL_Referee) ProtoMessage()    {}
func (*SSL_Referee) Descriptor() ([]byte, []int) {
	return fileDescriptor_ssl_referee_7314bdac51db39bf, []int{0}
}
func (m *SSL_Referee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSL_Referee.Unmarshal(m, b)
}
func (m *SSL_Referee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSL_Referee.Marshal(b, m, deterministic)
}
func (dst *SSL_Referee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSL_Referee.Merge(dst, src)
}
func (m *SSL_Referee) XXX_Size() int {
	return xxx_messageInfo_SSL_Referee.Size(m)
}
func (m *SSL_Referee) XXX_DiscardUnknown() {
	xxx_messageInfo_SSL_Referee.DiscardUnknown(m)
}

var xxx_messageInfo_SSL_Referee proto.InternalMessageInfo

func (m *SSL_Referee) GetPacketTimestamp() uint64 {
	if m != nil && m.PacketTimestamp != nil {
		return *m.PacketTimestamp
	}
	return 0
}

func (m *SSL_Referee) GetStage() SSL_Referee_Stage {
	if m != nil && m.Stage != nil {
		return *m.Stage
	}
	return SSL_Referee_NORMAL_FIRST_HALF_PRE
}

func (m *SSL_Referee) GetStageTimeLeft() int32 {
	if m != nil && m.StageTimeLeft != nil {
		return *m.StageTimeLeft
	}
	return 0
}

func (m *SSL_Referee) GetCommand() SSL_Referee_Command {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return SSL_Referee_HALT
}

func (m *SSL_Referee) GetCommandCounter() uint32 {
	if m != nil && m.CommandCounter != nil {
		return *m.CommandCounter
	}
	return 0
}

func (m *SSL_Referee) GetCommandTimestamp() uint64 {
	if m != nil && m.CommandTimestamp != nil {
		return *m.CommandTimestamp
	}
	return 0
}

func (m *SSL_Referee) GetYellow() *SSL_Referee_TeamInfo {
	if m != nil {
		return m.Yellow
	}
	return nil
}

func (m *SSL_Referee) GetBlue() *SSL_Referee_TeamInfo {
	if m != nil {
		return m.Blue
	}
	return nil
}

func (m *SSL_Referee) GetDesignatedPosition() *SSL_Referee_Point {
	if m != nil {
		return m.DesignatedPosition
	}
	return nil
}

func (m *SSL_Referee) GetBlueTeamOnPositiveHalf() bool {
	if m != nil && m.BlueTeamOnPositiveHalf != nil {
		return *m.BlueTeamOnPositiveHalf
	}
	return false
}

func (m *SSL_Referee) GetGameEvent() *SSL_Referee_Game_Event {
	if m != nil {
		return m.GameEvent
	}
	return nil
}

// Information about a single team.
type SSL_Referee_TeamInfo struct {
	// The team's name (empty string if operator has not typed anything).
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// The number of goals scored by the team during normal play and overtime.
	Score *uint32 `protobuf:"varint,2,req,name=score" json:"score,omitempty"`
	// The number of red cards issued to the team since the beginning of the game.
	RedCards *uint32 `protobuf:"varint,3,req,name=red_cards,json=redCards" json:"red_cards,omitempty"`
	// The amount of time (in microseconds) left on each yellow card issued to the team.
	// If no yellow cards are issued, this array has no elements.
	// Otherwise, times are ordered from smallest to largest.
	YellowCardTimes []uint32 `protobuf:"varint,4,rep,packed,name=yellow_card_times,json=yellowCardTimes" json:"yellow_card_times,omitempty"`
	// The total number of yellow cards ever issued to the team.
	YellowCards *uint32 `protobuf:"varint,5,req,name=yellow_cards,json=yellowCards" json:"yellow_cards,omitempty"`
	// The number of timeouts this team can still call.
	// If in a timeout right now, that timeout is excluded.
	Timeouts *uint32 `protobuf:"varint,6,req,name=timeouts" json:"timeouts,omitempty"`
	// The number of microseconds of timeout this team can use.
	TimeoutTime *uint32 `protobuf:"varint,7,req,name=timeout_time,json=timeoutTime" json:"timeout_time,omitempty"`
	// The pattern number of this team's goalie.
	Goalie *uint32 `protobuf:"varint,8,req,name=goalie" json:"goalie,omitempty"`
	// The total number of bot collisions committed by this team
	BotCollisions *uint32 `protobuf:"varint,9,opt,name=bot_collisions,json=botCollisions" json:"bot_collisions,omitempty"`
	// The number of consecutive ball placement failures of this team
	BallPlacementFailures *uint32 `protobuf:"varint,10,opt,name=ball_placement_failures,json=ballPlacementFailures" json:"ball_placement_failures,omitempty"`
	// The total number of infringements where robots were too fast during game stoppage
	BotSpeedInfringements *uint32  `protobuf:"varint,11,opt,name=bot_speed_infringements,json=botSpeedInfringements" json:"bot_speed_infringements,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *SSL_Referee_TeamInfo) Reset()         { *m = SSL_Referee_TeamInfo{} }
func (m *SSL_Referee_TeamInfo) String() string { return proto.CompactTextString(m) }
func (*SSL_Referee_TeamInfo) ProtoMessage()    {}
func (*SSL_Referee_TeamInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ssl_referee_7314bdac51db39bf, []int{0, 0}
}
func (m *SSL_Referee_TeamInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSL_Referee_TeamInfo.Unmarshal(m, b)
}
func (m *SSL_Referee_TeamInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSL_Referee_TeamInfo.Marshal(b, m, deterministic)
}
func (dst *SSL_Referee_TeamInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSL_Referee_TeamInfo.Merge(dst, src)
}
func (m *SSL_Referee_TeamInfo) XXX_Size() int {
	return xxx_messageInfo_SSL_Referee_TeamInfo.Size(m)
}
func (m *SSL_Referee_TeamInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SSL_Referee_TeamInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SSL_Referee_TeamInfo proto.InternalMessageInfo

func (m *SSL_Referee_TeamInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SSL_Referee_TeamInfo) GetScore() uint32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *SSL_Referee_TeamInfo) GetRedCards() uint32 {
	if m != nil && m.RedCards != nil {
		return *m.RedCards
	}
	return 0
}

func (m *SSL_Referee_TeamInfo) GetYellowCardTimes() []uint32 {
	if m != nil {
		return m.YellowCardTimes
	}
	return nil
}

func (m *SSL_Referee_TeamInfo) GetYellowCards() uint32 {
	if m != nil && m.YellowCards != nil {
		return *m.YellowCards
	}
	return 0
}

func (m *SSL_Referee_TeamInfo) GetTimeouts() uint32 {
	if m != nil && m.Timeouts != nil {
		return *m.Timeouts
	}
	return 0
}

func (m *SSL_Referee_TeamInfo) GetTimeoutTime() uint32 {
	if m != nil && m.TimeoutTime != nil {
		return *m.TimeoutTime
	}
	return 0
}

func (m *SSL_Referee_TeamInfo) GetGoalie() uint32 {
	if m != nil && m.Goalie != nil {
		return *m.Goalie
	}
	return 0
}

func (m *SSL_Referee_TeamInfo) GetBotCollisions() uint32 {
	if m != nil && m.BotCollisions != nil {
		return *m.BotCollisions
	}
	return 0
}

func (m *SSL_Referee_TeamInfo) GetBallPlacementFailures() uint32 {
	if m != nil && m.BallPlacementFailures != nil {
		return *m.BallPlacementFailures
	}
	return 0
}

func (m *SSL_Referee_TeamInfo) GetBotSpeedInfringements() uint32 {
	if m != nil && m.BotSpeedInfringements != nil {
		return *m.BotSpeedInfringements
	}
	return 0
}

// The coordinates of the Designated Position. These are measured in
// millimetres and correspond to SSL-Vision coordinates. These fields are
// always either both present (in the case of a ball placement command) or
// both absent (in the case of any other command).
type SSL_Referee_Point struct {
	X                    *float32 `protobuf:"fixed32,1,req,name=x" json:"x,omitempty"`
	Y                    *float32 `protobuf:"fixed32,2,req,name=y" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSL_Referee_Point) Reset()         { *m = SSL_Referee_Point{} }
func (m *SSL_Referee_Point) String() string { return proto.CompactTextString(m) }
func (*SSL_Referee_Point) ProtoMessage()    {}
func (*SSL_Referee_Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_ssl_referee_7314bdac51db39bf, []int{0, 1}
}
func (m *SSL_Referee_Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSL_Referee_Point.Unmarshal(m, b)
}
func (m *SSL_Referee_Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSL_Referee_Point.Marshal(b, m, deterministic)
}
func (dst *SSL_Referee_Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSL_Referee_Point.Merge(dst, src)
}
func (m *SSL_Referee_Point) XXX_Size() int {
	return xxx_messageInfo_SSL_Referee_Point.Size(m)
}
func (m *SSL_Referee_Point) XXX_DiscardUnknown() {
	xxx_messageInfo_SSL_Referee_Point.DiscardUnknown(m)
}

var xxx_messageInfo_SSL_Referee_Point proto.InternalMessageInfo

func (m *SSL_Referee_Point) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *SSL_Referee_Point) GetY() float32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*SSL_Referee)(nil), "SSL_Referee")
	proto.RegisterType((*SSL_Referee_TeamInfo)(nil), "SSL_Referee.TeamInfo")
	proto.RegisterType((*SSL_Referee_Point)(nil), "SSL_Referee.Point")
	proto.RegisterEnum("SSL_Referee_Stage", SSL_Referee_Stage_name, SSL_Referee_Stage_value)
	proto.RegisterEnum("SSL_Referee_Command", SSL_Referee_Command_name, SSL_Referee_Command_value)
}

func init() { proto.RegisterFile("ssl_referee.proto", fileDescriptor_ssl_referee_7314bdac51db39bf) }

var fileDescriptor_ssl_referee_7314bdac51db39bf = []byte{
	// 892 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0x5e, 0x9b, 0xff, 0x03, 0x86, 0xe1, 0x84, 0x24, 0x2e, 0xbd, 0xa1, 0xa9, 0xda, 0xb2, 0xaa,
	0xca, 0xc5, 0x4a, 0xdd, 0x7b, 0xc7, 0x6b, 0x12, 0x14, 0x07, 0xa3, 0xc1, 0xab, 0x76, 0xaf, 0x46,
	0x0e, 0x0c, 0xc8, 0xaa, 0xb1, 0x91, 0x3d, 0xd9, 0x6e, 0x5e, 0xa0, 0x0f, 0xd0, 0x77, 0xe9, 0x83,
	0xf4, 0x8d, 0xaa, 0x19, 0xdb, 0x81, 0x90, 0xb6, 0x77, 0x33, 0xdf, 0xf7, 0x9d, 0x6f, 0xce, 0x39,
	0x9f, 0x01, 0xfa, 0x59, 0x16, 0xb1, 0x94, 0x6f, 0x78, 0xca, 0xf9, 0x64, 0x9f, 0x26, 0x22, 0x19,
	0x0e, 0x24, 0xb4, 0x0d, 0x76, 0x9c, 0xf1, 0xcf, 0x3c, 0x16, 0x39, 0x7a, 0xf5, 0x97, 0x01, 0xed,
	0xe5, 0xd2, 0x65, 0x34, 0xd7, 0xe2, 0x5b, 0x20, 0xfb, 0x60, 0xf5, 0x1b, 0x17, 0x4c, 0x84, 0x3b,
	0x9e, 0x89, 0x60, 0xb7, 0x37, 0xb5, 0x91, 0x3e, 0xae, 0xd2, 0x5e, 0x8e, 0xfb, 0x25, 0x8c, 0x63,
	0xa8, 0x65, 0x22, 0xd8, 0x72, 0x53, 0x1f, 0xe9, 0xe3, 0xee, 0x3b, 0x9c, 0x1c, 0xf9, 0x4c, 0x96,
	0x92, 0xa1, 0xb9, 0x00, 0xbf, 0x87, 0x9e, 0x3a, 0x28, 0x4f, 0x16, 0xf1, 0x8d, 0x30, 0x2b, 0x23,
	0x6d, 0xdc, 0xa7, 0x86, 0x82, 0xa5, 0xa5, 0xcb, 0x37, 0x02, 0x27, 0xd0, 0x58, 0x25, 0xbb, 0x5d,
	0x10, 0xaf, 0xcd, 0xaa, 0xf2, 0x1c, 0xbc, 0xf0, 0xb4, 0x73, 0x8e, 0x96, 0x22, 0xfc, 0x01, 0x7a,
	0xc5, 0x91, 0xad, 0x92, 0xc7, 0x58, 0xf0, 0xd4, 0xac, 0x8d, 0xf4, 0xb1, 0x41, 0xbb, 0x05, 0x6c,
	0xe7, 0x28, 0xfe, 0x08, 0xfd, 0x52, 0x78, 0x18, 0xab, 0xae, 0xc6, 0x22, 0x05, 0x71, 0x98, 0xeb,
	0x27, 0xa8, 0x3f, 0xf1, 0x28, 0x4a, 0x7e, 0x37, 0x1b, 0x23, 0x7d, 0xdc, 0x7e, 0x77, 0xfe, 0xa2,
	0x09, 0x9f, 0x07, 0xbb, 0x59, 0xbc, 0x49, 0x68, 0x21, 0xc2, 0xb7, 0x50, 0x7d, 0x88, 0x1e, 0xb9,
	0xd9, 0xfc, 0x3f, 0xb1, 0x92, 0xa0, 0x0d, 0x67, 0x6b, 0x9e, 0x85, 0xdb, 0x38, 0x10, 0x7c, 0xcd,
	0xf6, 0x49, 0x16, 0x8a, 0x30, 0x89, 0xcd, 0xd6, 0x48, 0x1b, 0xb7, 0x4f, 0xf6, 0xb7, 0x48, 0xc2,
	0x58, 0x50, 0x3c, 0xc8, 0x17, 0x85, 0x1a, 0xdf, 0xc3, 0x85, 0x34, 0x93, 0xd6, 0x5e, 0x9c, 0xa3,
	0x9f, 0xf9, 0x6d, 0x10, 0x6d, 0x4c, 0x18, 0x69, 0xe3, 0x26, 0xfd, 0x0f, 0x16, 0x7f, 0x86, 0x96,
	0x4c, 0xdf, 0x91, 0xe1, 0x9b, 0x6d, 0xf5, 0xe4, 0xe5, 0xf1, 0x93, 0xec, 0x46, 0x7e, 0x1b, 0x8a,
	0xa6, 0x07, 0xe5, 0xf0, 0xcf, 0x0a, 0x34, 0xcb, 0x31, 0x10, 0xa1, 0x1a, 0x07, 0x3b, 0xae, 0xbe,
	0x88, 0x16, 0x55, 0x67, 0x1c, 0x40, 0x2d, 0x5b, 0x25, 0x69, 0xfe, 0x19, 0x18, 0x34, 0xbf, 0xe0,
	0xd7, 0xd0, 0x4a, 0xf9, 0x9a, 0xad, 0x82, 0x74, 0x9d, 0x99, 0x15, 0xc5, 0x34, 0x53, 0xbe, 0xb6,
	0xe5, 0x1d, 0x27, 0xd0, 0xcf, 0x97, 0xa7, 0xf8, 0x3c, 0x12, 0xb3, 0x3a, 0xaa, 0x8c, 0x8d, 0x6b,
	0x9d, 0x68, 0xb4, 0x97, 0x93, 0x52, 0xab, 0x52, 0xc1, 0x6f, 0xa0, 0x73, 0xa4, 0xcf, 0x8a, 0x90,
	0xdb, 0x07, 0x59, 0x86, 0x43, 0x68, 0x4a, 0x9b, 0xe4, 0x51, 0x64, 0x2a, 0x58, 0x83, 0x3e, 0xdf,
	0x65, 0x79, 0x71, 0x56, 0x4f, 0xa9, 0x58, 0x0d, 0xda, 0x2e, 0x30, 0xf9, 0x04, 0x5e, 0x40, 0x7d,
	0x9b, 0x04, 0x51, 0x98, 0xc7, 0x68, 0xd0, 0xe2, 0x86, 0xdf, 0x41, 0xf7, 0x21, 0x11, 0x6c, 0x95,
	0x44, 0x51, 0x98, 0x85, 0x49, 0x9c, 0xa9, 0xb0, 0x0c, 0x6a, 0x3c, 0x24, 0xc2, 0x7e, 0x06, 0xf1,
	0x3d, 0x5c, 0x3e, 0x04, 0x51, 0xc4, 0xf6, 0x51, 0xb0, 0xe2, 0x3b, 0x1e, 0x0b, 0xb6, 0x09, 0xc2,
	0xe8, 0x31, 0xe5, 0x99, 0x0a, 0xc5, 0xa0, 0xe7, 0x92, 0x5e, 0x94, 0xec, 0xb4, 0x20, 0x55, 0x5d,
	0x22, 0x58, 0xb6, 0xe7, 0x7c, 0xcd, 0xc2, 0x78, 0x93, 0x86, 0xf1, 0x56, 0x29, 0x32, 0x95, 0x90,
	0xac, 0x4b, 0xc4, 0x52, 0xb2, 0xb3, 0x63, 0x72, 0xf8, 0x2d, 0xd4, 0xd4, 0x07, 0x82, 0x1d, 0xd0,
	0xbe, 0xa8, 0x34, 0x74, 0xaa, 0x7d, 0x91, 0xb7, 0x27, 0x15, 0x83, 0x4e, 0xb5, 0xa7, 0xab, 0xbf,
	0x75, 0xa8, 0xa9, 0x9f, 0x21, 0x7e, 0x05, 0xe7, 0x73, 0x8f, 0xde, 0x5b, 0x2e, 0x9b, 0xce, 0xe8,
	0xd2, 0x67, 0xb7, 0x96, 0x3b, 0x65, 0x0b, 0xea, 0x90, 0x37, 0x78, 0x0e, 0xfd, 0x57, 0x14, 0xd1,
	0x70, 0x00, 0xa4, 0x80, 0x95, 0xd6, 0x9f, 0xdd, 0x3b, 0x44, 0xc7, 0x21, 0x5c, 0x14, 0xe8, 0xd2,
	0xb1, 0xbd, 0xf9, 0x87, 0x83, 0x51, 0x05, 0x2f, 0x00, 0x5f, 0x73, 0xa4, 0x2a, 0x9d, 0x9c, 0x5f,
	0x7d, 0x6a, 0x29, 0x0f, 0x76, 0x4d, 0x1d, 0xeb, 0x8e, 0xd4, 0xd0, 0x84, 0x41, 0x8e, 0x9e, 0x34,
	0x54, 0x3f, 0xe8, 0x8f, 0xfa, 0x69, 0xe0, 0x19, 0xf4, 0x72, 0xf4, 0xd0, 0x4e, 0x53, 0x8e, 0x95,
	0x83, 0xa7, 0xdd, 0xb4, 0xe4, 0x58, 0xaf, 0x28, 0x02, 0x72, 0x80, 0x85, 0x33, 0xb7, 0x5c, 0xff,
	0x13, 0x5b, 0xde, 0x7a, 0x9e, 0xef, 0x7d, 0xf4, 0x8b, 0x96, 0xda, 0xf2, 0xe1, 0x53, 0x8e, 0x74,
	0xd0, 0x80, 0xd6, 0xc2, 0x5b, 0xfa, 0xec, 0xc6, 0xba, 0x77, 0x88, 0x71, 0xf5, 0x47, 0x05, 0x1a,
	0xc5, 0xdf, 0x10, 0x36, 0xa1, 0x7a, 0x6b, 0xb9, 0x3e, 0x79, 0x23, 0x4f, 0x4b, 0xdf, 0x5b, 0x10,
	0x0d, 0x09, 0x74, 0xca, 0x2d, 0xf8, 0x16, 0xf5, 0x89, 0x8e, 0x3d, 0x68, 0x4f, 0x3d, 0x6a, 0x3b,
	0x05, 0x50, 0x51, 0x3d, 0x50, 0x67, 0x61, 0x51, 0x87, 0xdd, 0xcd, 0xec, 0x3b, 0x6f, 0x3a, 0x65,
	0x9f, 0x1c, 0xd7, 0xf5, 0x7e, 0x21, 0x55, 0xb9, 0x96, 0x53, 0xee, 0xda, 0xfd, 0xe8, 0x90, 0xda,
	0x71, 0x55, 0xd9, 0x65, 0x51, 0x55, 0x3f, 0xae, 0x2a, 0x39, 0x55, 0xd5, 0x90, 0xa1, 0x7c, 0x98,
	0x51, 0xc7, 0xf6, 0xd9, 0x94, 0x3a, 0x4e, 0x59, 0xd1, 0x94, 0xb3, 0x1e, 0xe3, 0x4a, 0xdd, 0x92,
	0x3e, 0xb3, 0xf9, 0xbf, 0xe8, 0x41, 0xfa, 0xbc, 0x64, 0x54, 0x45, 0x1b, 0x11, 0xba, 0x32, 0x0b,
	0xb9, 0xc6, 0x42, 0xdb, 0x91, 0x2b, 0x28, 0x31, 0xa5, 0x32, 0xe4, 0x0a, 0x6e, 0x3c, 0xcb, 0x2d,
	0x25, 0x5d, 0xb9, 0x54, 0x05, 0x28, 0xbe, 0x27, 0x73, 0xbc, 0xb6, 0x5c, 0x97, 0x2d, 0x5c, 0xcb,
	0x76, 0xee, 0x9d, 0xf9, 0xb3, 0x19, 0xc1, 0x4b, 0x38, 0x3b, 0xa1, 0x54, 0x4d, 0xff, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbb, 0x38, 0xc2, 0xbc, 0xe1, 0x06, 0x00, 0x00,
}
