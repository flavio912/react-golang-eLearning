import * as React from 'react';

export type IconNames =
  | 'AddDelegateRepeater'
  | 'ArrowLeft'
  | 'ArrowLeftBlue'
  | 'ArrowLeftNavyBlue'
  | 'ArrowRight'
  | 'ArrowRightNavyBlue'
  | 'AviationSecurityCert'
  | 'Basket'
  | 'Blue_TTC_Logo_Icon'
  | 'Card_SecondaryActon_Dots'
  | 'CloseCourseManagementTray_X'
  | 'CloseHamburger'
  | 'CourseAccountActivated'
  | 'CourseCertificates'
  | 'CourseExpiringSoon'
  | 'CourseFailed'
  | 'CourseNewCourse'
  | 'CourseNewCourseGreen'
  | 'CourseNewCourseGrey'
  | 'CourseStatus_Completed'
  | 'CourseStatus_Incomplete'
  | 'CourseStatus_NotStarted'
  | 'CourseTimeTracked'
  | 'CourseTimeTrackedGreen'
  | 'CourseTimeTrackedGrey'
  | 'Course_Calendar'
  | 'DB_Schenker'
  | 'Dhl'
  | 'Dhl_Grey'
  | 'Down_Arrow'
  | 'DownloadCSV'
  | 'EditDelegateProfilePic_Default'
  | 'EditPicPencil'
  | 'Facebook_Logo'
  | 'FilterAdjust'
  | 'FormCheckbox_Checked'
  | 'FormCheckbox_states'
  | 'FormCheckbox_Unchecked'
  | 'Heathrow'
  | 'Icon_Delegates'
  | 'LeftNav_Icon_Courses'
  | 'LeftNav_Icon_Dashboard'
  | 'LeftNav_Icon_Delegates'
  | 'LeftNav_Icon_Training'
  | 'LinkedIn_Logo'
  | 'Loading_Screen_Donut'
  | 'Location_Pin'
  | 'PDF_Icon'
  | 'RemoveSelectedCourse_X'
  | 'Right_Arrow'
  | 'SampleImage_ClassroomCoursesDetail_Feat'
  | 'SearchGlass'
  | 'TTC_Logo_Icon'
  | 'TTC_Logo'
  | 'Maersk'
  | 'Maersk_Grey'
  | 'MenuHamburger'
  | 'Mp3_Pause'
  | 'Mp3_Speaker'
  | 'Mp3_Play'
  | 'Nippon_Express'
  | 'Payments_Method'
  | 'Stripe'
  | 'Tick'
  | 'Twitter_Logo'
  | 'Ups'
  | 'Ups_Grey'
  | 'Volume'
  | 'Arrow_Right_Blue';

const iconNameMap = {
  AddDelegateRepeater: require('../../../assets/AddDelegateRepeater.svg'),
  ArrowLeft: require('../../../assets/ArrowLeft.svg'),
  ArrowLeftBlue: require('../../../assets/ArrowLeftBlue.svg'),
  ArrowLeftNavyBlue: require('../../../assets/ArrowLeftNavyBlue.svg'),
  ArrowRight: require('../../../assets/ArrowRight.svg'),
  ArrowRightNavyBlue: require('../../../assets/ArrowRightNavyBlue.svg'),
  AviationSecurityCert: require('../../../assets/AviationSecurityCert.svg'),
  Basket: require('../../../assets/Basket.svg'),
  Blue_TTC_Logo_Icon: require('../../../assets/logo/blue-ttc-logo.svg'),
  Card_SecondaryActon_Dots: require('../../../assets/Card_SecondaryActon_Dots.svg'),
  CloseCourseManagementTray_X: require('../../../assets/CloseCourseManagementTray_X.svg'),
  CloseHamburger: require('../../../assets/CloseHamburger.svg'),
  CourseAccountActivated: require('../../../assets/CourseAccountActivated.svg'),
  CourseCertificates: require('../../../assets/CourseCertificates.svg'),
  CourseExpiringSoon: require('../../../assets/CourseExpiringSoon.svg'),
  CourseFailed: require('../../../assets/CourseFailed.svg'),
  CourseNewCourse: require('../../../assets/CourseNewCourse.svg'),
  CourseNewCourseGreen: require('../../../assets/CourseNewCourseGreen.svg'),
  CourseNewCourseGrey: require('../../../assets/CourseNewCourseGrey.svg'),
  CourseStatus_Completed: require('../../../assets/CourseStatus_Completed.svg'),
  CourseStatus_Incomplete: require('../../../assets/CourseStatus_Incomplete.svg'),
  CourseStatus_NotStarted: require('../../../assets/CourseStatus_NotStarted.svg'),
  CourseTimeTracked: require('../../../assets/CourseTimeTracked.svg'),
  CourseTimeTrackedGreen: require('../../../assets/CourseTimeTrackedGreen.svg'),
  CourseTimeTrackedGrey: require('../../../assets/CourseTimeTrackedGrey.svg'),
  Course_Calendar: require('../../../assets/Course_Calendar.svg'),
  DB_Schenker: require('../../../assets/db-schenker-logo.png'),
  Dhl: require('../../../assets/dhl-logo.svg'),
  Dhl_Grey: require('../../../assets/dhl-logo-grey.png'),
  Down_Arrow: require('../../../assets/Down_Arrow.svg'),
  DownloadCSV: require('../../../assets/DownloadCSV.svg'),
  EditDelegateProfilePic_Default: require('../../../assets/EditDelegateProfilePic_Default.png'),
  EditPicPencil: require('../../../assets/EditPicPencil.svg'),
  Facebook_Logo: require('../../../assets/Facebook_Logo.png'),
  FilterAdjust: require('../../../assets/FilterAdjust.svg'),
  FormCheckbox_Checked: require('../../../assets/FormCheckbox_Checked.svg'),
  FormCheckbox_states: require('../../../assets/FormCheckbox_states.png'),
  FormCheckbox_Unchecked: require('../../../assets/FormCheckbox_Unchecked.svg'),
  Heathrow: require('../../../assets/heathrow-logo.png'),
  Icon_Delegates: require('../../../assets/Icon_Delegates.svg'),
  LeftNav_Icon_Courses: require('../../../assets/LeftNav_Icon_Courses.svg'),
  LeftNav_Icon_Dashboard: require('../../../assets/LeftNav_Icon_Dashboard.svg'),
  LeftNav_Icon_Delegates: require('../../../assets/LeftNav_Icon_Delegates.svg'),
  LeftNav_Icon_Training: require('../../../assets/LeftNav_Icon_Training.svg'),
  LinkedIn_Logo: require('../../../assets/LinkedIn_Logo.png'),
  Loading_Screen_Donut: require('../../../assets/Loading_Screen_Donut.svg'),
  Location_Pin: require('../../../assets/Location_Pin.svg'),
  PDF_Icon: require('../../../assets/PDF_Icon.svg'),
  RemoveSelectedCourse_X: require('../../../assets/RemoveSelectedCourse_X.svg'),
  Right_Arrow: require('../../../assets/Right_Arrow.svg'),
  SampleImage_ClassroomCoursesDetail_Feat: require('../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png'),
  SearchGlass: require('../../../assets/SearchGlass.svg'),
  TTC_Logo_Icon: require('../../../assets/logo/ttc-logo-icon.svg'),
  TTC_Logo: require('../../../assets/logo/ttc-logo.svg'),
  Maersk: require('../../../assets/maersk-logo.svg'),
  MenuHamburger: require('../../../assets/MenuHamburger.svg'),
  Maersk_Grey: require('../../../assets/maersk-logo-grey.png'),
  Mp3_Speaker: require('../../../assets/Mp3_Speaker.svg'),
  Mp3_Pause: require('../../../assets/Mp3_Pause.svg'),
  Nippon_Express: require('../../../assets/nippon-express-logo.png'),
  Payments_Method: require('../../../assets/Payments_Method.svg'),
  Stripe: require('../../../assets/Stripe.svg'),
  Mp3_Play: require('../../../assets/Mp3_Play.svg'),
  Tick: require('../../../assets/Tick.svg'),
  Twitter_Logo: require('../../../assets/Twitter_Logo.png'),
  Ups: require('../../../assets/ups-logo.svg'),
  Ups_Grey: require('../../../assets/ups-logo-grey.png'),
  Volume: require('../../../assets/Volume.svg'),
  Arrow_Right_Blue: require('../../../assets/Arrow_Right_Blue.svg')
};

type Props = {
  name: IconNames;
  pointer?: boolean;
  style?: React.CSSProperties;
  size?: number | null;
};

function Icon({
  name,
  size = 20,
  style,
  pointer,
  ...props
}: Props & React.ImgHTMLAttributes<HTMLImageElement>) {
  const cursor = pointer ? 'pointer' : 'auto';
  const sizeStyle = size ? { height: size, width: size } : {};
  return (
    <img
      style={{
        cursor,
        ...sizeStyle,
        ...style
      }}
      src={iconNameMap[name]}
      alt={`${name} icon`}
      {...props}
    />
  );
}

export default Icon;
